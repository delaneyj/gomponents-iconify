package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoSmallBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 8c4.97 0 9-.672 9-1.5c0-.52-1.588-.978-4-1.247M12 5c-4.97 0-9 .672-9 1.5c0 .589 2.037 1.099 5 1.344M20 10s-2.791 1-8.276 1C6.24 11 4 10 4 10m12 6s-1.258.5-4 .5s-4-.5-4-.5m5 3c-1.352 0-2.342-.121-3-.245m6-.255a9.627 9.627 0 0 1-1 .289M18 13s-.515 1-6 1m-6-1s.219.425 1.892.721"/>`),
		g.Group(children),
	)
}