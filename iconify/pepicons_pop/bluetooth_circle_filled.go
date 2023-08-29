package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopBluetoothCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m14.535 13.795l5.572-4.25a1 1 0 0 0 0-1.59l-5.572-4.25a1 1 0 0 0-1.606.795V13a1 1 0 0 0 1.606.795Zm.394-7.274L17.85 8.75l-2.922 2.23V6.52Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m14.535 22.295l5.572-4.25a1 1 0 0 0 0-1.59l-5.572-4.25A1 1 0 0 0 12.93 13v8.5a1 1 0 0 0 1.606.795Zm.394-7.274l2.922 2.229l-2.922 2.23v-4.46Z" clip-rule="evenodd"/><path d="M14.526 12.198a1 1 0 1 1-1.195 1.604L6.366 8.607a1 1 0 1 1 1.196-1.603l6.964 5.194Z"/><path d="M14.526 13.802a1 1 0 1 0-1.195-1.604l-6.965 5.195a1 1 0 0 0 1.196 1.603l6.964-5.194Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopBluetoothCircleFilled0)"/></g>`),
		g.Group(children),
	)
}