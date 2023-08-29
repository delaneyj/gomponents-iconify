package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NauseatedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#7cb342" d="M127.1 94.3c0 13.8-28.2 25-63 25s-63-11.2-63-25s1.5-83.6 63-83.6c61.6 0 63 69.8 63 83.6"/><path fill="#2f2f2f" d="M42 68c-4.5 0-8.2-4.3-8.2-9.6c0-5.4 3.6-9.7 8-9.8c4.5 0 8.2 4.3 8.2 9.6c.1 5.4-3.5 9.8-8 9.8zm44.6 0c-4.5 0-8.2-4.3-8.2-9.6c0-5.4 3.6-9.7 8-9.8c4.5 0 8.2 4.3 8.2 9.6c.1 5.4-3.5 9.8-8 9.8z"/><path fill="none" stroke="#2f2f2f" stroke-linecap="round" stroke-miterlimit="10" stroke-width="6" d="M40.4 79.8s3.3 1.3 5 6.2c1.7 4.8.2 10.2.2 10.2m42-16.4s-3.3 1.3-5 6.2c-1.7 4.8-.2 10.2-.2 10.2m-36.2-8.7c11.7-2.8 24-2.7 35.6.3"/><path fill="#7cb342" d="M79.9 48.6s.9 2.8 7.7 4.7c4.3 1.2 7.9.9 7.9.9s.1-7.5-7.9-8.1c-5.1-.5-7.7 2.5-7.7 2.5z"/>`),
		g.Group(children),
	)
}