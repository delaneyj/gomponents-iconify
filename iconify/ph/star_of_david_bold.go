package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfDavidBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m205.82 128l28.6-50A12 12 0 0 0 224 60h-57l-28.55-50A12 12 0 0 0 128 4a12 12 0 0 0-10.42 6.05L89 60H32a12 12 0 0 0-10.42 18l28.58 50l-28.58 50A12 12 0 0 0 32 196h57l28.57 50a12 12 0 0 0 20.84 0L167 196h57a12 12 0 0 0 10.42-17.95Zm-2.5-44L192 103.81L180.68 84Zm-25.14 44L153 172h-50.1l-25.09-44l25.13-44H153ZM128 40.18L139.33 60h-22.67ZM52.68 84h22.61L64 103.82Zm0 87.92L64 152.18L75.29 172Zm75.32 43.9L116.66 196h22.67ZM180.68 172L192 152.19L203.32 172Z"/>`),
		g.Group(children),
	)
}