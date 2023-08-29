package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsLr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLr0)"><path fill="#eee" d="M0 232.7L256 0h256l-19.2 24L512 46.4v46.7l-19.6 22.6l19.6 23.9v46.5L493.8 209l18.2 23.7v46.6L493 304l19 21.8v46.6l-18 23.5l18 23v46.6l-253.3 21L0 465.5v-46.6l18.8-21.6L0 372.4v-46.6l19.6-21.9L0 279.3z"/><path fill="#d80027" d="m256 0l-22 46.5h278V0zm-17.1 93.2v46.5H512V93.2zm-4.9 93l22 46.5h256v-46.5zM0 279.3v46.5h512v-46.5zm0 93.1v46.5h512v-46.5zm0 93.1V512h512v-46.5z"/><path fill="#0052b4" d="M0 0h256v232.7H0z"/><path fill="#eee" d="m152.4 66.8l16.6 51h53.6l-43.4 31.5l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.5h53.6z"/></g>`),
		g.Group(children),
	)
}