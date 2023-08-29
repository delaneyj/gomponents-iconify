package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.67 23.444L12 22.5l.33.944l-.33.116l-.33-.116Zm.33-2.016l.009-.004a17.755 17.755 0 0 0 3.887-2.215C18.034 17.59 20 15.223 20 12V5.36l-8-2.8l-8 2.8V12c0 3.223 1.966 5.588 4.104 7.21A17.75 17.75 0 0 0 12 21.427Zm-5.104-.625C4.534 19.012 2 16.127 2 12V3.94l10-3.5l10 3.5V12c0 4.127-2.534 7.012-4.896 8.803a19.75 19.75 0 0 1-4.65 2.595a6.99 6.99 0 0 1-.087.033l-.025.009l-.007.002l-.005.002L12 22.5l-.33.944l-.005-.002l-.007-.002l-.025-.01a12.04 12.04 0 0 1-.397-.155a19.756 19.756 0 0 1-4.34-2.473ZM11 16.5V10h2v6.5h-2Zm2-8h-2.004V6.496H13V8.5Z"/>`),
		g.Group(children),
	)
}