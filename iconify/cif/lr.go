package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><linearGradient id="cifLr0" x1="0%" x2="100%" y1="50%" y2="50%"><stop offset="0%" stop-color="#FFF"/><stop offset="100%"/></linearGradient></defs><g fill="none"><path fill="#BF0A30" d="M.5.552h300v157.895H.5z"/><path fill="url(#cifLr0)" stroke="#FFF" stroke-width="13.158" d="M.5 22.084h300M.5 50.792h300M.5 79.5h300M.5 108.208h300M.5 136.916h300"/><path fill="#002868" d="M.5.552h71.77v71.77H.5z"/><path fill="#FFF" d="m15.908 29.784l12.656 9.195l-4.834 14.878l12.655-9.195l12.656 9.195l-4.834-14.878l12.656-9.195H41.219l-4.834-14.877l-4.834 14.877z"/></g>`),
		g.Group(children),
	)
}