package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowtie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 493q0 26-13 46t-41 28t-72-15q-7-8-20.5-21T827 487t-71-51q-32-18-73-38q21-41 21-86q0-82-59-138q32-25 62-43q37-21 85-36t77-19l29-5q44-22 72-13.5t41 28.5t13 46q-7 10-17.5 28T978 226.5T960 316q0 40 16 84t32 69zm-512-53q-53 0-90.5-37.5T384 312t37.5-90.5T512 184t90.5 37.5T640 312t-37.5 90.5T512 440zm-171-42q-41 20-73 38q-34 20-69.5 49T145 533l-19 19q-44 23-72 15t-41-28t-13-46q7-10 17.5-27T46 402t18-86q0-41-16-87t-32-71L0 132q0-26 13-46t41-28.5T126 71q12 2 32 5t71 18.5t88 36.5q30 17 62 42q-59 57-59 139q0 45 21 86z"/>`),
		g.Group(children),
	)
}