package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSouthAfrica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="M60.9 40c.7-2.5 1.1-5.2 1.1-8s-.4-5.5-1.1-8H31.7L14.8 7.4c-4 2.8-7.3 6.6-9.6 11L19 32L5.3 45.6c2.2 4.4 5.5 8.2 9.6 11L31.7 40h29.2z"/><path fill="#3e4347" d="M3.5 22.5c-1 3-1.5 6.2-1.5 9.5s.5 6.5 1.5 9.5l9.6-9.5l-9.6-9.5"/><path fill="#428bc1" d="M18.3 58.7C22.4 60.8 27.1 62 32 62c12.3 0 22.9-7.4 27.5-18.1H33.4L18.3 58.7z"/><path fill="#ed4c5c" d="M59.5 20.1C54.9 9.4 44.3 2 32 2c-4.9 0-9.6 1.2-13.7 3.3l15.1 14.8h26.1"/><path fill="#fff" d="M60.5 22.7c-.3-.9-.6-1.8-1-2.6H33.4L18.3 5.3c-.7.3-1.3.7-2 1.1c-.5.3-1 .7-1.5 1L31.7 24h29.2c-.1-.4-.2-.9-.4-1.3"/><path fill="#ffce31" d="M5.3 18.4c-.1.2-.2.4-.3.7c-.5 1.1-1 2.2-1.4 3.4l9.6 9.5l-9.6 9.5c.4 1.2.8 2.3 1.4 3.4c.1.2.2.4.3.7L19 32L5.3 18.4"/><path fill="#fff" d="M31.7 40L14.8 56.6c.3.2.7.5 1 .7c.8.5 1.6 1 2.5 1.4l15.1-14.8h26.1c.4-1 .8-2 1.2-3.1c.1-.3.1-.5.2-.8H31.7"/>`),
		g.Group(children),
	)
}