package tndao

import (
    "fmt"
    "strconv"

    "github.com/rocket-pool/rocketpool-go/dao"
    "github.com/rocket-pool/rocketpool-go/types"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services/rocketpool"
    cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)


func voteOnProposal(c *cli.Context) error {

    // Get RP client
    rp, err := rocketpool.NewClientFromCtx(c)
    if err != nil { return err }
    defer rp.Close()

    // Get trusted node DAO proposals
    proposals, err := rp.TNDAOProposals()
    if err != nil {
        return err
    }

    // Get votable proposals
    votableProposals := []dao.ProposalDetails{}
    for _, proposal := range proposals.Proposals {
        if proposal.State == types.Active && !proposal.MemberVoted {
            votableProposals = append(votableProposals, proposal)
        }
    }

    // Check for votable proposals
    if len(votableProposals) == 0 {
        fmt.Println("No proposals can be voted on.")
        return nil
    }

    // Get selected proposal
    var selectedProposal dao.ProposalDetails
    if c.String("proposal") != "" {

        // Get selected proposal ID
        selectedId, err := strconv.ParseUint(c.String("proposal"), 10, 64)
        if err != nil {
            return fmt.Errorf("Invalid proposal ID '%s': %w", c.String("proposal"), err)
        }

        // Get matching proposal
        found := false
        for _, proposal := range votableProposals {
            if proposal.ID == selectedId {
                selectedProposal = proposal
                found = true
                break
            }
        }
        if !found {
            return fmt.Errorf("Proposal %d can not be voted on.", selectedId)
        }

    } else {

        // Prompt for proposal selection
        options := make([]string, len(votableProposals))
        for pi, proposal := range votableProposals {
            options[pi] = fmt.Sprintf(
                "proposal %d (message: '%s', payload: %s, end block: %d, votes required: %.2f, votes for: %.2f, votes against: %.2f)",
                proposal.ID,
                proposal.Message,
                proposal.PayloadStr,
                proposal.EndBlock,
                proposal.VotesRequired,
                proposal.VotesFor,
                proposal.VotesAgainst)
        }
        selected, _ := cliutils.Select("Please select a proposal to vote on:", options)
        selectedProposal = votableProposals[selected]

    }

    // Get support status
    var support bool
    if c.String("support") != "" {

        // Parse support status
        var err error
        support, err = cliutils.ValidateBool("support", c.String("support"))
        if err != nil { return err }

    } else {

        // Prompt for support status
        support = cliutils.Confirm("Would you like to vote in support of the proposal?")

    }

    // Vote on proposal
    if _, err := rp.VoteOnTNDAOProposal(selectedProposal.ID, support); err != nil {
        return err
    }

    // Log & return
    if support {
        fmt.Printf("Successfully voted in support of proposal %d.\n", selectedProposal.ID)
    } else {
        fmt.Printf("Successfully voted against proposal %d.\n", selectedProposal.ID)
    }
    return nil

}

