package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneBluetoothSpeakerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12V8.2l-2.3 2.3l-.7-.7L16.8 7L14 4.2l.7-.7L17 5.8V2h.5l2.85 2.9L18.2 7l2.15 2.15L17.5 12H17Zm1-1.9l.95-.95L18 8.2v1.9Zm0-4.3l.95-.9l-.95-.95V5.8ZM19.95 21q-3.125 0-6.188-1.35T8.2 15.8q-2.5-2.5-3.85-5.55T3 4.05V3h5.9l.925 5.025l-2.85 2.875q.55.975 1.225 1.85t1.45 1.625q.725.725 1.588 1.388T13.1 17l2.9-2.9l5 1.025V21h-1.05Z"/>`),
		g.Group(children),
	)
}