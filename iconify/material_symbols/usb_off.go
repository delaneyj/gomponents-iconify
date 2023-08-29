package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L13.15 16H13v2.3q.475.25.738.713T14 20q0 .825-.588 1.413T12 22q-.825 0-1.413-.588T10 20q0-.525.275-.975T11 18.3V16H8q-.825 0-1.413-.588T6 14v-2.3q-.475-.25-.738-.713T5 10q0-.425.163-.813t.512-.662l-4.3-4.3L2.8 2.8l18.4 18.4l-1.425 1.425Zm-2-7.7L16 13.15V12h-1V8h4v4h-1v2q0 .25-.05.488t-.175.437ZM8 14h3v-.15l-2.525-2.525q-.1.125-.225.213T8 11.7V14Zm5-3.85l-2-2V6H9l3-4l3 4h-2v4.15Z"/>`),
		g.Group(children),
	)
}