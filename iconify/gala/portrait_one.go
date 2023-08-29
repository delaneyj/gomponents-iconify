package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortraitOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaPortrait10" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaPortrait11" stroke-opacity="1" d="M 16.056766,240 H 239.94325"/><path id="galaPortrait12" stroke-opacity="1" d="m 16.056766,240 c 0,-47.97569 31.983762,-64.0813 31.983762,-64.0813 15.991919,-9.23291 31.9838,0 47.975681,-15.99189"/><path id="galaPortrait13" stroke-opacity="1" d="m 96.016209,159.92681 c 0,0 15.991921,15.99189 31.983801,15.99189"/><path id="galaPortrait14" stroke-opacity="1" d="m 239.94325,240 c 0,-47.97569 -31.9838,-64.0813 -31.9838,-64.0813 -15.99188,-9.23291 -31.98376,0 -47.97564,-15.99189"/><path id="galaPortrait15" stroke-opacity="1" d="m 159.98381,159.92681 c 0,0 -15.99192,15.99189 -31.9838,15.99189"/><path id="galaPortrait16" d="m 128.00001,15.99977 c 31.9838,0 47.97568,15.991881 47.97568,63.967561 0,39.979759 -31.9838,63.967599 -47.97568,63.967599"/><path id="galaPortrait17" d="m 128.00001,15.99977 c -31.983801,0 -47.975682,15.991881 -47.975682,63.967561 0,39.979759 31.983802,63.967599 47.975682,63.967599"/><path id="galaPortrait18" stroke-opacity="1" d="m 159.98381,159.92681 -7.41232,-27.66304"/><path id="galaPortrait19" stroke-opacity="1" d="m 96.016209,159.92681 7.412321,-27.66304"/></g>`),
		g.Group(children),
	)
}