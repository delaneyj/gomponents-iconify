package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motorola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12C24.002 5.375 18.632.002 12.007 0H12zm7.327 18.065s-.581-2.627-1.528-4.197c-.514-.857-1.308-1.553-2.368-1.532c-.745 0-1.399.423-2.2 1.553c-.469.77-.882 1.573-1.235 2.403c0 0-.29-.675-.63-1.343a8.038 8.038 0 0 0-.605-1.049c-.804-1.13-1.455-1.539-2.2-1.553c-1.049-.021-1.854.675-2.364 1.528c-.948 1.574-1.528 4.197-1.528 4.197h-.864l4.606-15.12l3.56 11.804l.024.021l.024-.021l3.56-11.804l4.61 15.113h-.862z"/>`),
		g.Group(children),
	)
}