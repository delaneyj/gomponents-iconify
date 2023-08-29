package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1536 1024l341 1024H0l341-1024h512V650q-57-14-104-45t-80-76t-53-97t-19-112q0-71 27-133t73-108T806 6t133-27q70 0 132 27t109 73t73 108t27 133q0 58-19 111t-52 98t-81 75t-104 46v374h512zM768 320q0 35 13 66t37 55t54 36t67 14q35 0 66-13t54-37t36-54t14-67q0-35-13-66t-37-54t-54-37t-66-14q-35 0-66 13t-55 37t-36 55t-14 66zm-304 875l-227 682h1403l-227-682h-389v256q0 35-25 60t-60 25q-18 0-33-6t-27-18t-19-27t-7-34v-256H464z"/>`),
		g.Group(children),
	)
}