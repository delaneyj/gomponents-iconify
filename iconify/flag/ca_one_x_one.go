package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M81.1 0h362.3v512H81.1z"/><path fill="#d52b1e" d="M-100 0H81.1v512H-100zm543.4 0h181.1v512H443.4zM135.3 247.4l-14 4.8l65.4 57.5c5 14.8-1.7 19.1-6 26.9l71-9l-1.8 71.5l14.8-.5l-3.3-70.9l71.2 8.4c-4.4-9.3-8.3-14.2-4.3-29l65.4-54.5l-11.4-4.1c-9.4-7.3 4-34.8 6-52.2c0 0-38.1 13.1-40.6 6.2l-9.9-18.5l-34.6 38c-3.8 1-5.4-.6-6.3-3.8l16-79.7l-25.4 14.3c-2.1.9-4.2 0-5.6-2.4l-24.5-49l-25.2 50.9c-1.9 1.8-3.8 2-5.4.8l-24.2-13.6l14.5 79.2c-1.1 3-3.9 4-7.1 2.3l-33.3-37.8c-4.3 7-7.3 18.4-13 21c-5.7 2.3-25-4.9-37.9-7.7c4.4 15.9 18.2 42.3 9.5 51z"/>`),
		g.Group(children),
	)
}