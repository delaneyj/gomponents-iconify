package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AgOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagAg1x10"><path fill="#25ff01" d="M109 47.6h464.8v464.9H109z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagAg1x10)" transform="translate(-120 -52.4) scale(1.1014)"><path fill="#fff" d="M0 47.6h693V512H0z"/><path d="M109 47.6h464.8v186.1H109V47.6Z"/><path fill="#0072c6" d="M128.3 232.1h435.8v103.5H128.3V232.1Z"/><path fill="#ce1126" d="M692.5 49.2v463.3H347L692.5 49.2zm-691.3 0v463.3h345.7L1.2 49.2z"/><path fill="#fcd116" d="m508.8 232.2l-69.3-17.6l59-44.4l-72.5 10.3l37.3-63l-64.1 37.2l11.3-73.5l-43.4 58l-17.6-67.3l-19.6 69.3l-43.4-59l12.4 75.6l-64.1-39.3l37.2 63l-70.3-11.3l57.9 43.4l-72.4 18.6h321.6z"/></g>`),
		g.Group(children),
	)
}