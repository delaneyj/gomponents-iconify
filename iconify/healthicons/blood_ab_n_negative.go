package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodAbNNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBloodAbNNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM20.388 5.535a5.018 5.018 0 0 1 7.224 0l.528.548h5.19c2.573 0 4.67 2.074 4.67 4.646V15.5h-2v-4.771a2.658 2.658 0 0 0-2.67-2.646h-5.616a1 1 0 0 1-.72-.307l-.823-.855a3.018 3.018 0 0 0-4.342 0l-.823.855a1 1 0 0 1-.72.307H15c-1.663 0-3 1.338-3 2.974V28.97a68.776 68.776 0 0 1 3.7.06l.53.02c1.284.048 2.481.093 3.652.069c2.627-.056 5.042-.466 7.61-1.981c3.118-1.84 5.758-1.288 7.583-.226c.338.197.647.41.925.626V22.5h2v10.526C38 35.78 35.755 38 33 38h-3v2h-5v4h-2v-4h-5v-2h-3c-2.755 0-5-2.22-5-4.974v-21.97c0-2.752 2.245-4.973 5-4.973h4.86l.528-.548Zm-.972 9.064a1 1 0 0 0-1.832 0l-2.616 5.98a.998.998 0 0 0-.018.041l-.866 1.98a1 1 0 0 0 1.832.8l.613-1.4h3.942l.613 1.4a1 1 0 0 0 1.832-.8l-.866-1.98a.998.998 0 0 0-.018-.04l-2.616-5.98Zm-.916 2.896L19.596 20h-2.192l1.096-2.505ZM24 15a1 1 0 0 1 1-1h3.125C29.773 14 31 15.405 31 17a3.09 3.09 0 0 1-.732 2c.46.54.732 1.249.732 2c0 1.595-1.227 3-2.875 3H25a1 1 0 0 1-1-1v-8Zm4.125 3c.423 0 .875-.386.875-1s-.452-1-.875-1H26v2h2.125ZM26 20h2.125c.423 0 .875.386.875 1s-.452 1-.875 1H26v-2Zm8-2a1 1 0 1 0 0 2h3a1 1 0 0 0 0-2h-3Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBloodAbNNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}