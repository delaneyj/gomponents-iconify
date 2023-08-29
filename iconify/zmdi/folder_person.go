package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 64q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h128l42 43h171zm-106.5 64q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5zM363 299v-22q0-19-29.5-30.5t-56-11.5t-56 11.5T192 277v22h171z"/>`),
		g.Group(children),
	)
}