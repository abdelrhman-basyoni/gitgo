package core

import "fmt"

type Commit struct {
	parent []byte

	tree []byte

	oid []byte

	message string

	author Author

	committer Author
}

func (c *Commit) New(parent, tree []byte, message string, author, committer Author) error {
	c.parent = parent
	c.tree = tree

	c.message = message
	c.author = author
	c.committer = committer
	return nil
}

func (c *Commit) ToString() string {

	parent := func() string {
		if c.parent == nil {
			return ""
		}

		return fmt.Sprintf("parent %x\n", c.parent)
	}

	return fmt.Sprintf("tree %x\n%sauthor %s\ncommitter %s\n\n%s", c.tree, parent(), c.author.ToString(), c.committer.ToString(), c.message)
}

func (c *Commit) GetTree() []byte {
	return c.tree
}

func (c *Commit) GetOid() []byte {
	return c.oid
}

func (c *Commit) SetOid(oid []byte) {
	c.oid = oid
}

func (c *Commit) GetType() string {
	return "commit"
}

func (c *Commit) GetParent() []byte {
	return c.parent
}

func (c *Commit) GetMessage() string {
	return c.message
}
