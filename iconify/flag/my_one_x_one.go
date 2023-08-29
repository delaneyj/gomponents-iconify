package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#C00" d="M0 0h512v36.6H0z"/><path fill="#fff" d="M0 36.6h512V73H0z"/><path fill="#C00" d="M0 73.1h512v36.6H0z"/><path fill="#fff" d="M0 109.7h512v36.6H0z"/><path fill="#C00" d="M0 146.3h512v36.6H0z"/><path fill="#fff" d="M0 182.9h512v36.5H0z"/><path fill="#C00" d="M0 219.4h512V256H0z"/><path fill="#fff" d="M0 256h512v36.6H0z"/><path fill="#C00" d="M0 292.6h512V329H0z"/><path fill="#fff" d="M0 329.1h512v36.6H0z"/><path fill="#C00" d="M0 365.7h512v36.6H0z"/><path fill="#fff" d="M0 402.3h512v36.6H0z"/><path fill="#C00" d="M0 438.9h512v36.5H0z"/><path fill="#fff" d="M0 475.4h512V512H0z"/><path fill="#006" d="M0 0h256v292.6H0V0Z"/><path fill="#FC0" d="m166 93l4.8 32.5l18.4-27.2l-10 31.3l28.5-16.6l-22.5 24l32.8-2.6l-30.7 11.9L218 158l-32.8-2.5l22.5 24l-28.4-16.7l9.8 31.5l-18.4-27.3l-4.8 32.5l-4.7-32.5l-18.4 27.2l9.9-31.4l-28.4 16.7l22.4-24l-32.8 2.5l30.7-11.8l-30.6-11.9l32.8 2.6l-22.5-24l28.4 16.6l-10-31.4l18.5 27.3l4.8-32.6Zm-26.7 1.3a56.9 56.9 0 0 0-73 24.9a56.9 56.9 0 0 0 45.5 83.8a56.9 56.9 0 0 0 27.5-4.7a64 64 0 1 1 0-104Z"/>`),
		g.Group(children),
	)
}