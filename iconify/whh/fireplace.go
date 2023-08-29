package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fireplace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 448h-32v544q0 13-9.5 22.5T928 1024H96q-13 0-22.5-9.5T64 992V448H32q-13 0-22.5-9.5T0 416V288q0-13 9.5-22.5T32 256h139l85-128V32q0-13 9.5-22.5T288 0h448q13 0 22.5 9.5T768 32v96l85 128h139q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5T992 448zM419 960h166q7-25 7-64q0-21-18-61t-36-70l-18-29q0 5-2 13t-9 30.5t-16 42t-25.5 41T432 896q-25 16-13 64zM256 448v512h99q-35-37-35-80q0-32 18-67t48-61q35-31 51-67.5t14-66t-9.5-54T425 526l-9-14q7 3 17.5 8t41.5 21.5t58.5 35.5t59.5 48t53 60q29 41 44.5 90.5T704 864q-2 40-49 96h113V448H256z"/>`),
		g.Group(children),
	)
}