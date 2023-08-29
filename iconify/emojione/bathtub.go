package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bathtub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#c6bbb3" d="M13.1 57.5c0 2.5-.9 4.5-2.9 4.5s-3.4-2.2-2.5-4.5C9 53.9 12.2 52 14.2 52c1.9 0-1.1 3-1.1 5.5m37.8 0c0 2.5.9 4.5 2.9 4.5s3.4-2.2 2.5-4.5C55 53.9 51.8 52 49.8 52c-1.9 0 1.1 3 1.1 5.5"/><path fill="#e2d8c7" d="M59.1 40.1c0 18.4-12 18.8-26.9 18.8S5.4 58.5 5.4 40.1h53.7"/><path fill="#c6bbb3" d="M5.4 41h53.7v1.2H5.4z"/><path fill="#ddd3ca" d="M59.4 41.4H5.1c-4.1 0-4.1-4.8 0-4.8h54.3c3.5 0 3.5 4.8 0 4.8"/><path fill="#a6aeb0" d="M33.8 14.8C28.8 9.6 32.1 2 32.1 2s-10.5 10.1-3.9 17c7.7 8 3.9 13 3.9 13s9.6-9 1.7-17.2m-15 4.7c-3.8-3.8-1.3-9.3-1.3-9.3s-7.9 7.3-2.9 12.4c5.8 5.8 2.9 9.5 2.9 9.5s7.2-6.6 1.3-12.6m29.9 0c-3.8-3.8-1.3-9.3-1.3-9.3s-7.9 7.3-2.9 12.4c5.8 5.8 2.9 9.5 2.9 9.5s7.2-6.6 1.3-12.6" opacity=".6"/>`),
		g.Group(children),
	)
}