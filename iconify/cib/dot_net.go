package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotNet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.224 10.089v11.667h1.365v-8.438a12.58 12.58 0 0 0-.068-1.599h.052c.099.255.224.5.37.729l6 9.302h1.672V10.089h-1.359v8.203c-.016.573.016 1.146.083 1.714h-.031a13.986 13.986 0 0 0-.474-.781L5.995 10.09zm12.417 0v11.667h6.203l.005-1.281h-4.813v-4.047h4.214v-1.229h-4.214v-3.875h4.521V10.09zm7.25 0v1.234h3.354v10.432h1.365V11.323h3.391v-1.234zm-23.021 10a.86.86 0 0 0-.609.276a.91.91 0 0 0-.26.641a.901.901 0 0 0 1.542.641c.172-.167.271-.401.271-.641s-.099-.474-.271-.641a.876.876 0 0 0-.641-.276H.871z"/>`),
		g.Group(children),
	)
}