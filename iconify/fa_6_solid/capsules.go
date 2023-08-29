package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capsules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 144c0-26.5 21.5-48 48-48s48 21.5 48 48v112H64V144zm-64 0v224c0 61.9 50.1 112 112 112s112-50.1 112-112V189.6c1.8 19.1 8.2 38 19.8 54.8l128.5 187.3c35.5 51.7 105.3 64.3 156 28.1s63-107.5 27.5-159.2L427.3 113.3c-35.5-51.8-105.4-64.3-156-28.1c-28 20-44.3 50.8-47.3 83V144c0-61.9-50.1-112-112-112S0 82.1 0 144zm296.6 64.2c-16-23.3-10-55.3 11.9-71c21.2-15.1 50.5-10.3 66 12.2l67 97.6l-79.9 56l-65-94.8zM491 407.7c-.8.6-1.6 1.1-2.4 1.6l4-2.8c-.5.4-1 .8-1.6 1.2z"/>`),
		g.Group(children),
	)
}