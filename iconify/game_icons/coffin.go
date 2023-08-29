package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m312.094 21.97l59.094 119.686h52.093L362.595 22l-50.5-.03zM173 22.687l-63.594 127.218l65.844 345.75l114.688.094L354.467 150L291.626 22.75L173 22.687zm16.063 76.28h88.78v18.688h-88.78V98.97zm0 36.22h88.78v18.687h-88.78v-18.688zm182.5 25.156L309 495.438l47.25.03l68.313-335.124h-53z"/>`),
		g.Group(children),
	)
}