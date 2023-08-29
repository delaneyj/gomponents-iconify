package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emptynotepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FFB636" d="M488.326 452.829H229.07c-3.352 0-6.068-3.941-6.068-8.802V67.973c0-4.861 2.717-8.802 6.068-8.802h259.369c3.29 0 5.957 3.868 5.957 8.64v376.215c-.001 4.862-2.718 8.803-6.07 8.803z"/><path fill="#FFD469" d="M408.821 452.829H184.018L15.164 281.973v-214a8.802 8.802 0 0 1 8.802-8.802h384.855v393.658z"/><path fill="#FFB636" d="M184.018 452.829L15.164 283.975h132.963c19.822 0 35.891 16.069 35.891 35.891v132.963z"/>`),
		g.Group(children),
	)
}