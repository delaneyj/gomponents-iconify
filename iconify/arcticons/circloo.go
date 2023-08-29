package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circloo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.792 20.44a3.618 3.618 0 1 0 3.623 3.617a3.627 3.627 0 0 0-3.623-3.618Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.212 26.874c1.982 1.713 2.275 4.24 1.654 6.683a7.595 7.595 0 0 1-5.317 5.247a7.13 7.13 0 0 1-6.946-1.462m18.89 5.148l5.904-6.856m-6.063-6.128l-2.675-.554l-6.135 13.538M5.499 22.813l14.268 5.202M14.8 15.052l11.005 4.36m2.22-2.687l5.406-2.053m8.545-8.512L26.523 21.613M5.974 6.208l15.165 15.29M5.987 41.795l15.152-15.151m20.683 15.351L26.407 26.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.093 8.364c7.15.108 16.893-.747 23.958-2.845m5.423 5.143A22.563 22.563 0 0 0 42.5 36.008M28.725 19.411c-.5 3.645.198 10.192 3.073 12.582"/><circle cx="16.756" cy="28.708" r=".75" fill="currentColor"/><circle cx="13.671" cy="29.107" r=".75" fill="currentColor"/><circle cx="10.586" cy="29.506" r=".75" fill="currentColor"/><circle cx="7.501" cy="29.906" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}