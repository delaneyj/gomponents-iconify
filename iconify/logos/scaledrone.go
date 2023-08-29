package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scaledrone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosScaledrone0" x1="75.094%" x2="26.264%" y1="93.381%" y2="8.971%"><stop offset="0%" stop-color="#FFF" stop-opacity="0"/><stop offset="47.52%" stop-color="#6466BB"/><stop offset="100%" stop-color="#FFF" stop-opacity="0"/></linearGradient></defs><circle cx="128" cy="128" r="128" fill="url(#logosScaledrone0)" opacity=".5"/><path fill="#6466BB" d="M205.6 83.2L238.9 64c-30-52-96.6-69.9-148.6-39.8c-52 30-69.9 96.6-39.8 148.6L17.1 192c30 52 96.6 69.9 148.6 39.8c52.1-30 69.9-96.6 39.9-148.6Z"/>`),
		g.Group(children),
	)
}