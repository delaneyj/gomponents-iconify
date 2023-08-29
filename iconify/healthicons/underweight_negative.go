package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnderweightNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUnderweightNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm24 13a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm9.476 12.732a1.5 1.5 0 1 1-2.952.536c-.078-.43-.158-.943-.246-1.51a68.647 68.647 0 0 0-.55-3.183c-.174-.859-.36-1.595-.55-2.112a3.943 3.943 0 0 0-.118-.29a25.08 25.08 0 0 0-1.188-.274L27 25l1.5 5.5v12a1.5 1.5 0 0 1-3 0v-12h-3v12a1.5 1.5 0 0 1-3 0v-12L21 25l-.867-6.066c-.39.08-.788.173-1.196.277a3.924 3.924 0 0 0-.117.285c-.189.512-.374 1.244-.549 2.097a68.162 68.162 0 0 0-.551 3.184c-.088.56-.167 1.067-.244 1.491a1.5 1.5 0 1 1-2.952-.536c.052-.285.12-.716.2-1.23c.161-1.02.373-2.361.608-3.51c.183-.895.404-1.805.674-2.535a5.24 5.24 0 0 1 .522-1.075c.191-.287.587-.793 1.267-.974c4.413-1.177 7.973-1.232 12.413-.036c.681.184 1.077.692 1.266.979c.22.332.388.716.521 1.079c.27.733.49 1.647.673 2.545c.234 1.148.445 2.493.606 3.515c.081.52.15.957.202 1.242Zm-4.521-6.75s.004.008.015.02a.099.099 0 0 0-.015-.02Zm-9.913.039a.196.196 0 0 0-.014.02a.109.109 0 0 0 .014-.02Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUnderweightNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}