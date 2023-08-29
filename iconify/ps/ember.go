package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ember(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M380 160q-3 11-11.5 20.5T354 194l-7 4q-27-68-62-117t-56-64L208 2q-2 22-17.5 43.5T160 78l-15 11l-41-21q3 16-1 32.5T94 126l-5 9q-38 25-51 67.5T30 276l6 30l-36-17q8 55 30.5 93T80 434.5t52.5 22T176 462l17-1q48 0 84.5-17t56-43.5t32-58.5t15.5-64t3-58.5t-2-43.5zM186 425q-39 0-65.5-19T85 360.5t-12-53t-1-45.5l2-18q2 7 7 12.5t9 7.5l4 2q16-41 37-70t34-38l13-9q1 13 10 26t18 19l9 7l25-13q-2 9 .5 19.5T246 224l3 5q23 15 30.5 40.5T284 313l-4 18l22-10q-4 33-18 56t-30 32t-31.5 13t-26.5 4z"/>`),
		g.Group(children),
	)
}