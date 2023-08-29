package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WsFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#ce1126" d="M0 0h640v480H0z"/><path fill="#002b7f" d="M0 0h320v240H0z"/><path fill="#fff" d="m180 229.3l-20.7-14l-19.9 14.1l6.5-24.9l-19-15.2l24.5-1.5l8.1-23.6l8.8 24l24 .7l-19 16.3zm-3.6-165.6L159.8 53l-16 10.4l4.4-20l-14.6-12.7l19.4-1.6l7.2-18.6l7.4 18.7l19.1 1.7L172 44.3zm-73 59.5l-16-11l-16.7 11l5.2-19.4L60.8 91L80 90l7-19l6.8 18.9l19.6 1.1l-15 12.5zM250 110l-15.4-10l-15 10l4.4-18.3l-14-11.8l18.3-1.5l6.3-17.2l7 17.4l17.7 1l-13.7 12.3zm-43.1 43.4l-10.3-6.4l-10.3 6.6l2.7-12.3l-9.2-8.3l12-1l4.6-11.6l4.9 11.6l11.9 1l-9.1 8.3z"/></g>`),
		g.Group(children),
	)
}