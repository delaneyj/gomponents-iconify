package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buddhaquotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.87 40.175a10.765 10.765 0 0 1 3.108-1.931c-1.082-.48-2.276-.773-2.58-2.547a5.535 5.535 0 0 1 .404-2.6a26.71 26.71 0 0 1 1.597-7.901c1.482-4.936 7.49-4.165 9.054-8.32m-.001 0l-.751-.678a5.541 5.541 0 0 0-.7.476c-.018 1.869-2.702 1.409-1.44-1.982a3.94 3.94 0 0 0-.144-2.117c.005-1.597-.157-3.36.541-4.262c-.188-1.93.833-2.266 1.41-3.186m17.797 35.018a10.304 10.304 0 0 0-2.993-1.901c1.082-.48 2.275-.773 2.58-2.547a5.534 5.534 0 0 0-.405-2.6a26.707 26.707 0 0 0-1.596-7.901c-1.482-4.936-7.49-4.165-9.054-8.32m0 0l.752-.678a5.54 5.54 0 0 1 .699.476c.019 1.869 2.702 1.409 1.44-1.982a3.94 3.94 0 0 1 .144-2.117c-.004-1.597.158-3.36-.54-4.262c.187-1.93-.834-2.266-1.41-3.186"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.517 2.53c.024 0 .08.001.141.006a21.5 21.5 0 1 1-3.302-.001a1.842 1.842 0 0 1 .275 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.368 5.127c.345-1.879 1.399-1.948 2.248-2.539c.247-.932.695-1.547 1.459-1.664m3.706 4.203c-.345-1.879-1.398-1.948-2.248-2.539c-.247-.932-.695-1.547-1.458-1.664"/>`),
		g.Group(children),
	)
}