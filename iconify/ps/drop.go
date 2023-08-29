package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m229 59l-9-17q-7-16-23.5-26.5T163 5q-39 0-58 37l-8 17q-2 5-7 15Q8 256 3 308q0 69 47 118t113 49q65 0 112.5-49T323 308q0-11-2.5-25t-9.5-33t-12.5-34t-16.5-40t-16.5-38t-19-41.5T229 59zm-66 373q-48 0-83-36.5T45 308q0-42 90-230l8-17q8-13 20-13q11 0 19 13l8 17q90 182 90 230q0 52-34.5 88T163 432zm10-85q-21 0-32-13q-10-13-10-49q2-8-4-16t-13-8q-8-1-16.5 5.5T88 283q0 56 21 79q25 27 64 27q22 0 22-21t-22-21z"/>`),
		g.Group(children),
	)
}