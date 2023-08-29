package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#e30a17" d="M0 0h640v480H0z"/><path fill="#fff" d="M407 247.5c0 66.2-54.6 119.9-122 119.9s-122-53.7-122-120s54.6-119.8 122-119.8s122 53.7 122 119.9z"/><path fill="#e30a17" d="M413 247.5c0 53-43.6 95.9-97.5 95.9s-97.6-43-97.6-96s43.7-95.8 97.6-95.8s97.6 42.9 97.6 95.9z"/><path fill="#fff" d="m430.7 191.5l-1 44.3l-41.3 11.2l40.8 14.5l-1 40.7l26.5-31.8l40.2 14l-23.2-34.1l28.3-33.9l-43.5 12l-25.8-37z"/></g>`),
		g.Group(children),
	)
}