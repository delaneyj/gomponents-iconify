package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1408 928v320q0 119-84.5 203.5T1120 1536H288q-119 0-203.5-84.5T0 1248V416q0-119 84.5-203.5T288 128h704q14 0 23 9t9 23v64q0 14-9 23t-23 9H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V928q0-14 9-23t23-9h64q14 0 23 9t9 23zm384-864v512q0 26-19 45t-45 19t-45-19l-176-176l-652 652q-10 10-23 10t-23-10L695 983q-10-10-10-23t10-23l652-652l-176-176q-19-19-19-45t19-45t45-19h512q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}