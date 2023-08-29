package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 1)"><ellipse cx="5.967" cy="5.979" rx="1.967" ry="1.979"/><path d="M7.598.001c.002.035.014.068.014.105C7.61.788 6.887 1.34 5.994 1.34c-.892 0-1.615-.555-1.615-1.238c0-.035.01-.067.014-.102C1.864.714.003 3.072.001 5.88C-.003 9.259 2.678 11.999 5.985 12c3.31.003 5.994-2.732 5.996-6.108c.001-2.81-1.856-5.172-4.383-5.891zM2 5.997a4 4 0 1 1 8 .008a4 4 0 1 1-8-.008z"/></g><path d="M7.958 14.25a7.281 7.281 0 0 1-3.917-1.142l-1.824 1.811a.73.73 0 0 0 0 1.039l11.531-.052a.733.733 0 0 0 .002-1.039l-1.807-1.8a7.297 7.297 0 0 1-3.985 1.183z"/></g>`),
		g.Group(children),
	)
}