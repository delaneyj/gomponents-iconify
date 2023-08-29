package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 512q-25 0-57-23.5T648.5 425T597 322.5T576 192q0-16 1-24.5t6.5-19.5t19.5-15.5t37-4.5q43 0 94.5-20T824 64t76.5-44T960 0q29 0 46.5 17t17.5 47q0 73-20.5 144.5T952 333t-66.5 94.5t-66.5 63t-51 21.5zm-64 64q0 22-27.5 52t-64 57.5t-85 69T448 832q-31 37-52.5 85.5t-30 77.5t-13.5 29q-29 0-73.5-28t-92-79t-90-113.5T27 663T0 512q0-87 47-160.5T175 235t177-43q22 0 34 .5t34.5 2T456 200t27.5 11t21.5 18.5t7 26.5q0 55 26 114t70 110q13 15 38.5 35.5T688 551t16 25z"/>`),
		g.Group(children),
	)
}