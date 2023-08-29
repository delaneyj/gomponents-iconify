package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M149.275 512H0V362.725h149.275V512zm0-330.638H0v149.276h149.275V181.362zm181.363 181.363H181.362V512h149.276V362.725zm0-181.363H181.362v149.276h149.276V181.362zm0-181.362H181.362v149.275h149.276V0zM512 0H362.725v149.275H512V0zm0 181.362H362.725v149.276H512V181.362z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}