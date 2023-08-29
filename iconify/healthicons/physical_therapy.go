package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhysicalTherapy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="healthiconsPhysicalTherapy0" d="M23.916 27h6.078h-.04c-2.044-.048-3.635-.408-4.981-1.62a7.656 7.656 0 0 1-.993-1.104l-1.247 1.557L23.916 27Z"/><path id="healthiconsPhysicalTherapy1" fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm18.5 11a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm-7.252 12h-2.935l3.147 3.018l-.693 2.335c-.05.171-.124.335-.218.486L17.9 37.5a2 2 0 0 1-3.4-2.11l1.509-2.43l1.025-3.46l.127-.5H14a2 2 0 0 0-2 2v7h-2v-7a4 4 0 0 1 4-4h3.645a2.039 2.039 0 0 1-.09-.177c-.952-2.107-.277-4.354.482-5.823a9.575 9.575 0 0 1 1.422-2.055c.446-.485 1.142-1.124 1.99-1.367a7.89 7.89 0 0 1 .931-.226c.337-.053.893-.097 1.494.153c.598.249.953.67 1.14.923c.155.21.31.465.424.657l.04.066l.462.772c.332.557.602 1.008.863 1.41c.364.558.621.872.846 1.074c.318.286.772.556 2.397.593a2 2 0 0 1-.038 4H34a4 4 0 0 1 4 4v7h-2v-7a2 2 0 0 0-2-2h-8.057l1.254 1.238a2 2 0 0 1 .43.627l2.075 4.784a2 2 0 0 1-3.67 1.591l-1.92-4.428L20.248 29Z" clip-rule="evenodd"/></defs><g fill="currentColor"><use href="#healthiconsPhysicalTherapy0"/><use href="#healthiconsPhysicalTherapy1" fill-rule="evenodd" clip-rule="evenodd"/><use href="#healthiconsPhysicalTherapy0"/><use href="#healthiconsPhysicalTherapy1" fill-rule="evenodd" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}