package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GetPocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1565 0q65 0 110 45.5t45 110.5v519q0 176-68 336t-182.5 275t-274 182.5T861 1536q-176 0-335.5-67.5T251 1286T68 1011T0 675V156Q0 92 46 46T156 0h1409zM861 1064q47 0 82-33l404-388q37-35 37-85q0-49-34.5-83.5T1266 440q-47 0-82 33L861 783L538 473q-35-33-81-33q-49 0-83.5 34.5T339 558q0 51 36 85l405 388q33 33 81 33z"/>`),
		g.Group(children),
	)
}