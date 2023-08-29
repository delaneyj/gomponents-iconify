package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleChromeLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 8a92.08 92.08 0 0 1 80.78 48H128a44.05 44.05 0 0 0-43.82 40.11L54.51 72.72A91.9 91.9 0 0 1 128 36Zm0 128a36 36 0 1 1 36-36a36 36 0 0 1-36 36Zm-92-36a91.52 91.52 0 0 1 13.51-47.95L89.9 150c0 .09.11.17.17.26a43.93 43.93 0 0 0 56.47 17.63l-29.7 51.43A92.13 92.13 0 0 1 36 128Zm92 92h-2.29l40.39-70a1.21 1.21 0 0 0 .09-.2A43.89 43.89 0 0 0 153.25 92h59.41A92 92 0 0 1 128 220Z"/>`),
		g.Group(children),
	)
}