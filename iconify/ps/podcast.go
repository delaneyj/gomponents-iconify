package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 250v-9q0-99 67.5-169T232 2t162.5 70T462 241v9q-4-96-70.5-163T232 20T72.5 87T2 250zm230-143q58 0 100 41t46 101v-9q0-62-42.5-106.5T232 89t-103.5 44.5T86 240v9q4-60 46-101t100-41zm0 44q-19 0-33.5 15T184 201t14.5 35t33.5 15t33.5-15t14.5-35t-14.5-35t-33.5-15zm0 104q-18 0-27.5 1.5t-21 8.5t-16.5 23t-5 42q0 55 20.5 93.5T232 462t49.5-38.5T302 330q0-26-5-42t-16.5-23t-21-8.5T232 255z"/>`),
		g.Group(children),
	)
}