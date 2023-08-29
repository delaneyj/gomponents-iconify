package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDetectionLoudSoundSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.825 21.85L2.15 16.2l4.225-4.225V6.3l11.35 11.325H12.05L7.825 21.85Zm11.775-5.1l-1.45-1.45q1.125-2.05.788-4.338T16.95 7.026q-1.65-1.65-3.938-1.988t-4.337.788l-1.45-1.45q2.675-1.7 5.788-1.363T18.375 5.6q2.25 2.25 2.588 5.363T19.6 16.75Zm-2.95-2.95l-1.7-1.7q0-.625-.187-1.212t-.613-1.013q-.45-.45-1.038-.65t-1.237-.2l-1.7-1.7Q11.6 6.9 13.05 7.15t2.5 1.3q1.05 1.05 1.288 2.488T16.65 13.8Z"/>`),
		g.Group(children),
	)
}