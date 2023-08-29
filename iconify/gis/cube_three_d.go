package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m50 1.18l-.72 1.25l-6.973 12.062H47.5v42.98l-37.223 21.49l-2.597-4.498L0 87.783h1.441l13.934.008l-2.598-4.498L50 61.803l37.223 21.49l-2.598 4.498L100 87.783l-.72-1.25l-6.96-12.069l-2.597 4.499L52.5 57.473V14.491h5.193z" color="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M49.058 28.451L19.812 47.63c-.674.487-.855 1.041-.867 1.545v28.66c0 .63.328 1.216.867 1.545l29.246 19.178c.705.385 1.422.32 1.898-.01L80.188 79.38a1.81 1.81 0 0 0 .867-1.545v-28.66c-.006-.764-.33-1.167-.895-1.564l-29.218-19.16c-.696-.397-1.242-.208-1.884 0zm.942 3.66l25.515 16.906l-25.531 17.206L24.23 49.174ZM22.56 52.394L48.192 69.37v24.425l-25.63-16.976Zm54.88 0v24.424L51.808 93.794V69.368Z" color="currentColor"/>`),
		g.Group(children),
	)
}