package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 64H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24Zm-8 352H48V96h416Z"/><path fill="currentColor" d="M184 344a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.41 55.41 0 0 1 184 312a56 56 0 0 1 0-112a55.41 55.41 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 184 168a88 88 0 0 0 0 176Zm163.429 0a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.414 55.414 0 0 1 347.429 312a56 56 0 0 1 0-112a55.414 55.414 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 347.429 168a88 88 0 0 0 0 176Z"/>`),
		g.Group(children),
	)
}