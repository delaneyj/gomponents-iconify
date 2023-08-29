package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BestsignApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 27.584V9.146a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v18.438m0 0h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.53 23.646a18.159 18.159 0 0 1 1.86-5.757a17.102 17.102 0 0 1 5.291-6.257m-1.115 11.21a18.483 18.483 0 0 1 1.727-4.953a18.758 18.758 0 0 1 2.498-3.721m.663 8.321a18.615 18.615 0 0 1 1.655-4.6a19.438 19.438 0 0 1 1.117-1.905m6.631 17.395v5.604a1.874 1.874 0 0 1-1.868 1.868h0a1.567 1.567 0 0 1-1.308-.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.989 33.379h0a1.874 1.874 0 0 1 1.868 1.868v1.214a1.874 1.874 0 0 1-1.868 1.868h0a1.873 1.873 0 0 1-1.868-1.868v-1.214a1.873 1.873 0 0 1 1.868-1.868Zm-11.04 4.109a2.227 2.227 0 0 0 1.868.841h1.12a1.873 1.873 0 0 0 1.868-1.868h0a1.873 1.873 0 0 0-1.868-1.868h-1.214a1.873 1.873 0 0 1-1.868-1.868h0a1.873 1.873 0 0 1 1.868-1.868h1.121a2.006 2.006 0 0 1 1.868.84m-12.859 6.258a2.208 2.208 0 0 0 1.495.374h.373a1.236 1.236 0 0 0 1.214-1.214h0a1.236 1.236 0 0 0-1.214-1.214h-.84a1.236 1.236 0 0 1-1.215-1.215h0a1.236 1.236 0 0 1 1.215-1.214h.373c.84 0 1.214.094 1.495.374m2.475-2.055v6.538m-1.027-4.95h1.961m-8.572 4.016a1.808 1.808 0 0 1-1.588.934h0a1.873 1.873 0 0 1-1.868-1.868v-1.214a1.873 1.873 0 0 1 1.868-1.868h0a1.874 1.874 0 0 1 1.868 1.868v.654H11.13m32.37 2.355v-3.083a1.873 1.873 0 0 0-1.868-1.868h0a1.873 1.873 0 0 0-1.868 1.868v3.083m0-3.083v-1.868M7.582 34.593a1.868 1.868 0 1 1 0 3.736H4.5v-7.472h3.082a1.868 1.868 0 0 1 0 3.736Zm0 0H4.5"/><circle cx="32.008" cy="31.137" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.008 33.379v4.95"/>`),
		g.Group(children),
	)
}