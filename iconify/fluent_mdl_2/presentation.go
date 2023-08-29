package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Presentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h1920v128h-128v896q0 26-10 49t-27 41t-41 28t-50 10h-640v640h512v128H384v-128h512v-640H256q-26 0-49-10t-41-27t-28-41t-10-50V128H0V0zm1664 1024V128H256v896h1408zm-256-512v128H512V512h896z"/>`),
		g.Group(children),
	)
}