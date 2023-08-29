package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#3FB6D3" d="M12.3 64.2L76.3 0h39.4L32.1 83.6zm64 63.8h39.4L81.6 93.9l34.1-34.8H76.3L42.2 93.5z"/><path fill="#27AACD" d="m81.6 93.9l-20-20l-19.4 19.6l19.4 19.6z"/><path fill="#19599A" d="M115.7 128L81.6 93.9l-20 19.2L76.3 128z"/><linearGradient id="deviconFlutter0" x1="59.365" x2="86.825" y1="116.36" y2="99.399" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#1b4e94"/><stop offset=".63" stop-color="#1a5497"/><stop offset="1" stop-color="#195a9b"/></linearGradient><path fill="url(#deviconFlutter0)" d="m61.6 113.1l30.8-8.4l-10.8-10.8z"/>`),
		g.Group(children),
	)
}