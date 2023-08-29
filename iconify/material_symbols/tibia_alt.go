package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TibiaAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.85 22q-.4-1.175-.6-2.3T9 17.55q-.05-1.025-.012-1.875t.112-1.45q0-.025 0 0q-.55-.125-1.262-.312T6.313 13.4q-.813-.325-1.663-.787T2.975 11.5L9.05 3.025l6.425 2.925q1.6.725 2.563 2.2T19 11.425V22H9.85Zm3.65-3.025q.225 0 .425-.1t.325-.25q.15.15.35.25t.4.1q.425 0 .713-.288t.287-.712q0-.425-.288-.713T15 16.976V13q.425 0 .713-.288T16 12q0-.425-.288-.713T15 11q-.2 0-.4.088t-.35.237q-.125-.15-.325-.237T13.5 11q-.425 0-.713.288T12.5 12q0 .425.288.713T13.5 13v3.975q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Z"/>`),
		g.Group(children),
	)
}