package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackPainNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBackPainNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm32 30.833a3 3 0 1 1 6 0V43a1 1 0 1 1-2 0V30.833a1 1 0 1 0-2 0v.223a1 1 0 1 1-2 0v-.223ZM26 31l-.001-.063c.075-3.021.745-6.062 1.639-8.668c.214-.625.438-1.217.666-1.77l.757 2.992a2 2 0 0 0 1.196 1.366l5 2a2 2 0 0 0 1.486-3.714l-4.054-1.622l-1.75-6.918A2 2 0 0 0 29 13.093c-3.674 0-6.253 2.026-7.821 4.593C19.649 20.191 19 23.3 19 26c0 .934.33 1.843.65 2.54c.34.737.775 1.457 1.18 2.067c.41.617.816 1.159 1.118 1.545l.052.067V42a2 2 0 1 0 4 0V31ZM11.47 9.662a1 1 0 0 0-.94 1.766l3.715 1.974l-1.671.7a1 1 0 0 0-.097 1.798l4.937 2.724a1 1 0 0 0 .966-1.75l-3.126-1.726l1.71-.716a1 1 0 0 0 .083-1.805l-5.578-2.965ZM9.831 19.978c.11-.56.623-.927 1.145-.819l3.666.76a.968.968 0 0 1 .73.717c.096.368-.007.767-.269 1.036l-.797.82l1.938.504c.515.134.824.692.689 1.247c-.135.555-.663.896-1.179.762l-3.625-.944a.975.975 0 0 1-.69-.73c-.087-.363.018-.75.273-1.013l.724-.744l-1.862-.385c-.521-.108-.854-.65-.743-1.21ZM33 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBackPainNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}