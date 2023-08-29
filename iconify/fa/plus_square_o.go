package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusSquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 672v64q0 14-9 23t-23 9H768v352q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V768H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h352V288q0-14 9-23t23-9h64q14 0 23 9t9 23v352h352q14 0 23 9t9 23zm128 448V288q0-66-47-113t-113-47H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113zm128-832v832q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q119 0 203.5 84.5T1408 288z"/>`),
		g.Group(children),
	)
}