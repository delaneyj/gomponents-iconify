package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Violin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M937.06 929q-96-115-104-207l-18-18l16-16l-47-48l-144 144l48 48l16-16l9 9q93 14 211 116q-55 50-131 70.5t-148.5 6.5t-118.5-60q-19-19-39-51q25-34 25-75q0-53-37.5-90.5t-90.5-37.5q-19 0-40 7q-9-7-11-10q-60-59-73.5-156.5t31.5-178.5l-154-188q-15 14-36 14t-36-15l-50-50q-15-16-15-37.5t15-36.5l37-37q16-16 37.5-16t36.5 16l51 50q14 14 14.5 35t-13.5 36l188 154q80-45 177.5-31.5t157.5 73.5q2 3 10 11q-7 21-7 40q0 53 37 90.5t90 37.5q42 0 76-25q33 20 52 39q46 47 60 121t-8 151t-74 131zm-505-577l-80 80l80 80l80-80zm48 304l80 80l48-48l-80-80zm176-176l-48 48l80 80l48-48z"/>`),
		g.Group(children),
	)
}