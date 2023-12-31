package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 14 14")
)

func ComputerBatteryEmptyOnePhoneMobileChargeDeviceElectricityEmptyPowerBattery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 5.5A.5.5 0 0 0 13 5h-1V4a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1H11a1 1 0 0 0 1-1V9h1a.5.5 0 0 0 .5-.5Z"/>`), g.Group(children),
	)
}

func ComputerBatteryFullOnePhoneMobileChargeDeviceElectricityPowerBatteryFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 5.5A.5.5 0 0 0 13 5h-1V4a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1H11a1 1 0 0 0 1-1V9h1a.5.5 0 0 0 .5-.5Zm-10.25 0v3m3-3v3m3-3v3"/>`), g.Group(children),
	)
}

func ComputerBatteryLowOnePhoneMobileChargeDeviceElectricityPowerBatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 5.5A.5.5 0 0 0 13 5h-1V4a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1H11a1 1 0 0 0 1-1V9h1a.5.5 0 0 0 .5-.5Zm-10.25 0v3"/>`), g.Group(children),
	)
}

func ComputerBatteryMediumOnePhoneMobileChargeMediumDeviceElectricityPowerBattery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 5.5A.5.5 0 0 0 13 5h-1V4a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1H11a1 1 0 0 0 1-1V9h1a.5.5 0 0 0 .5-.5Zm-10.25 0v3m3-3v3"/>`), g.Group(children),
	)
}

func ComputerChipOneComputerDeviceChipElectronicsCpuMicroprocessor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="8" x="3" y="3" rx="1"/><path d="M5 3V.5M9 3V.5M3 9H.5M3 5H.5M9 11v2.5M5 11v2.5M11 5h2.5M11 9h2.5m-5-1.5h-2"/></g>`), g.Group(children),
	)
}

func ComputerChipTwoCoreMicroprocessorDeviceElectronicsChipComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="13" x="3" y=".5" rx="1"/><path d="M3 3.5H.5M3 7H.5M3 10.5H.5m13-7H11M13.5 7H11m2.5 3.5H11m-4.5 0h2"/></g>`), g.Group(children),
	)
}

func ComputerConnectionBluetoothBluetoothInternetServerNetworkWireless(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.25 9.5l7.5-5.5l-4-3.5v13l4-3.5l-7.5-5.5"/>`), g.Group(children),
	)
}

func ComputerConnectionCableSplitCablesCableSplitDeviceComputerElectronicsCordsCordSplitter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.88" r="1.5"/><path d="M5.5 13.38h3m-3-2.5h3m.249-7.754L11.252.623l2 2l-2.502 2.504zM2.75.63l-2 2m2 2l2-2"/></g>`), g.Group(children),
	)
}

func ComputerConnectionNetworkNetworkServerInternetEthernet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 10.5v3m-5 0h10"/><circle cx="7" cy="5.5" r="5"/><path d="M2 5.5h10m-5 5c3-3.42 3-6.76 0-10c-2.94 3.12-3 6.44 0 10Z"/></g>`), g.Group(children),
	)
}

func ComputerConnectionSignalLoadingBracketLoadingInternetAngleSignalServerNetworkConnecting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9.5L.5 7L3 4.5m8 5L13.5 7L11 4.5"/><circle cx="9" cy="7" r=".5"/><circle cx="5" cy="7" r=".5"/></g>`), g.Group(children),
	)
}

func ComputerConnectionUsbCableCablesCableDeviceCordComputerElectronicsCordsUsb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V11M4.5 4V1.5a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1V4m1 0h-7a.5.5 0 0 0-.5.5v2.59a1 1 0 0 0 .29.7L4.5 9v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V9l1.21-1.21a1 1 0 0 0 .29-.7V4.5a.5.5 0 0 0-.5-.5Zm-4 2.5h1"/>`), g.Group(children),
	)
}

func ComputerConnectionUsbPortCablesCableDeviceComputerPortElectronicsCordsCordUsb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.97" cy="11.5" r="2"/><path d="M6.97 9.5v-9M5.47 2L6.97.5L8.47 2M6 9.78a4.14 4.14 0 0 1-2.6-2.29"/><circle cx="3.22" cy="6.25" r="1.25"/><path d="M7 8.6a5.6 5.6 0 0 0 3.49-3.07"/><circle cx="10.78" cy="4.32" r="1.25"/></g>`), g.Group(children),
	)
}

func ComputerConnectionWifiWirelessWifiInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="10.94" r="1.31"/><path d="M4.53 8a3.49 3.49 0 0 1 5 0M2.36 6.31a6.55 6.55 0 0 1 9.29 0"/><path d="M.5 4.45a9.19 9.19 0 0 1 13 0"/></g>`), g.Group(children),
	)
}

func ComputerControllerOneRemoteQuadcopterDronesFlyingDroneControlControllerTechnologyFly(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="1"/><path d="M7 5.5v-5"/><circle cx="10" cy="9.5" r="1.5"/><circle cx="4" cy="9.5" r="1.5"/></g>`), g.Group(children),
	)
}

func ComputerDatabaseRaidStorageCodeDiskProgrammingDatabaseArrayHardDisc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="3" rx="6.5" ry="2.5"/><path d="M.5 3v8c0 1.38 2.91 2.5 6.5 2.5s6.5-1.12 6.5-2.5V3"/><path d="M13.5 7c0 1.38-2.91 2.5-6.5 2.5S.5 8.38.5 7"/></g>`), g.Group(children),
	)
}

func ComputerDatabaseServerOneServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="1" rx="1"/><circle cx="3.5" cy="4" r=".5"/><path d="M7.5 4H11M1.5 7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1"/><circle cx="3.5" cy="10" r=".5"/><path d="M7.5 10H11"/></g>`), g.Group(children),
	)
}

func ComputerDatabaseServerThreeServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="3" rx="1"/><circle cx="3.5" cy="6" r=".5"/><path d="M6.5 6h2M7 9v4.5m-5 0h10"/></g>`), g.Group(children),
	)
}

func ComputerDatabaseServerTwoServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="13" height="5" x=".5" y=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="currentColor" d="M3.25 2.25A.75.75 0 1 0 4 3a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 3h2M1.5 5.5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1"/><path fill="currentColor" d="M3.25 7.25A.75.75 0 1 0 4 8a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 8h2m-1 2.5v3m-5 0h10"/>`), g.Group(children),
	)
}

func ComputerDesktopAddDesktopDeviceDisplayAddPlusComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6m-.5-7h-5M7 4v5"/></g>`), g.Group(children),
	)
}

func ComputerDesktopBlockDesktopDeviceDisplayDisablePermissionComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><circle cx="7" cy="5" r="3"/><path d="m4.88 7.12l4.24-4.24"/></g>`), g.Group(children),
	)
}

func ComputerDesktopCheckSuccessApproveDeviceDisplayDesktopComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6M4.5 7l2 1.5l3.5-4"/></g>`), g.Group(children),
	)
}

func ComputerDesktopDeleteDeviceRemoveDisplayComputerDenyDesktopFailFailure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6m-1-5l-4-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ComputerDesktopFavoriteHeartDeviceDisplayComputerFavoriteLikeHeartDesktop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6M13 2a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2"/><path d="M7 4c1.17-2.59 3.5-1.29 3.5.65C10.5 7.2 7 8.5 7 8.5S3.5 7.2 3.5 4.6C3.5 2.66 5.83 1.36 7 4Z"/></g>`), g.Group(children),
	)
}

func ComputerDesktopFavoriteStarComputerDesktopDeviceDisplayLikeFavoriteStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><path d="M6.45 2.36a.59.59 0 0 1 1.1 0L8.19 4H9.9a.61.61 0 0 1 .56.39a.59.59 0 0 1-.16.61L8.79 6.34l.64 1.28a.58.58 0 0 1-.12.69a.59.59 0 0 1-.7.1L7 7.53l-1.61.88a.59.59 0 0 1-.7-.1a.58.58 0 0 1-.12-.69l.64-1.28L3.7 5a.59.59 0 0 1-.16-.65A.61.61 0 0 1 4.1 4h1.71Z"/></g>`), g.Group(children),
	)
}

func ComputerDesktopHelpDeviceHelpInformationDisplayComputerDesktopQuestionInfo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 2h1a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1m4 9l-1 2.5M8 11l1 2.5m-5 0h6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 2.5a2 2 0 1 1 2 2v1"/><path fill="currentColor" d="M7 7.5a.75.75 0 1 0 .75.75A.75.75 0 0 0 7 7.5Z"/>`), g.Group(children),
	)
}

func ComputerDesktopLockDeviceSecureDisplayComputerLockDesktopPadlockSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h2.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><path d="M5 5.5h4v3H5zm.5 0v-1a1.5 1.5 0 0 1 3 0v1"/></g>`), g.Group(children),
	)
}

func ComputerDesktopRemoveDesktopDeviceSubtractDisplayMinusComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6m-.5-7h-5"/></g>`), g.Group(children),
	)
}

func ComputerDesktopScreenDeviceDisplayComputerDesktopElectronicsMonitorKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="8" x="1.5" y=".5" rx="1"/><path d="M12.28 11.05a1 1 0 0 0-.9-.55H2.62a1 1 0 0 0-.9.55l-.5 1a1 1 0 0 0 .05 1a1 1 0 0 0 .85.47h9.76a1 1 0 0 0 .85-.47a1 1 0 0 0 0-1ZM1.5 6h11"/></g>`), g.Group(children),
	)
}

func ComputerDesktopSettingSettingsDesktopDisplayDeviceGearCogComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2h1a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1m4 9l-1 2.5M8 11l1 2.5m-5 0h6M7.03 2v1.5M4 3.75l1.3.75M4 7.25l1.3-.75M7.03 9V7.5m3.03-.25l-1.3-.75m1.3-2.75l-1.3.75"/><circle cx="7.03" cy="5.5" r="2"/></g>`), g.Group(children),
	)
}

func ComputerDesktopWarningAlertDesktopDisplayDeviceWarningNotificationComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 2H1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-8A.5.5 0 0 0 13 2h-2.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6m-3-12V5"/><circle cx="7" cy="8" r=".5"/></g>`), g.Group(children),
	)
}

func ComputerHandHeldTabletKindleDeviceElectronicsIpadComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="M1.5 10.5h11M4.5 3h5m-5 2.5h5M4.5 8h3"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardAltWindowsKeyboardKeyAltPc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2.5v-1a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v1m0 9v1a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-.75M10.5 5h3M12 5v4M6.5 5v4h2m-5 0V6.5a1.5 1.5 0 0 0-3 0V9m0-1.5h3"/>`), g.Group(children),
	)
}

func ComputerKeyboardAsteriskOneAsteriskStarKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v13m-5.5-10l11 7m-11 0l11-7"/>`), g.Group(children),
	)
}

func ComputerKeyboardAsteriskTwoAsteriskStarKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7H.5m10-5.5l-7 11m0-11l7 11"/>`), g.Group(children),
	)
}

func ComputerKeyboardCommandKeyboardMacCommandApple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.75 10A1.75 1.75 0 1 1 10 11.75v-9.5A1.75 1.75 0 1 1 11.75 4h-9.5A1.75 1.75 0 1 1 4 2.25v9.5A1.75 1.75 0 1 1 2.25 10Z"/>`), g.Group(children),
	)
}

func ComputerKeyboardComputerKeyboardDeviceElectronicsDvorakQwerty(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="1"/><path d="M3.5 10.75h7m-7-2.5h7m-7-2.75v-1a2 2 0 0 1 2-2h4a2 2 0 0 0 2-2"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardDeleteDeleteBackspaceKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M10.5 9h-5l-2-2l2-2h5v4z"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardEjectEjectKeyboardUnmountDismountRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M3.5 10h7m-7-2.5h7L7 3.5l-3.5 4z"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardNextKeyboardNextArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7h10M7 10.5L10.5 7L7 3.5m6.5 0v7"/>`), g.Group(children),
	)
}

func ComputerKeyboardOptionKeyboardOptionMacApple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M3.5 6h2L8 9h2.5m-2-3h2"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardPowerOffPowerKeyboardButtonOnOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5.5 3.5h-2v2m0-2L7 7m1.57-3.13a3.5 3.5 0 1 1-4.7 4.7"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardPreviousOnePreviousKeyboardArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7h-10M7 3.5L3.5 7L7 10.5m-6.5-7v7"/>`), g.Group(children),
	)
}

func ComputerKeyboardPreviousTwoPreviousKeyboardArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M8 5L6 7l2 2M6 7h4.5m-7-2.5v5"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardReturnOneKeyboardArrowReturnEnter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.5h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-4"/><path d="m3.5 7.5l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardReturnThreeEnterReturnKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="m5.5 10.5l-2-2l2-2"/><path d="M3.5 8.5h5a1 1 0 0 0 1-1v-3"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardReturnTwoKeyboardArrowReturnEnter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 9.5h9a4 4 0 0 0 0-8h-3"/><path d="m3.5 6.5l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardShiftKeyboardKeyShiftUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M8.5 10.5v-3h2L7 3.5l-3.5 4h2v3h3z"/></g>`), g.Group(children),
	)
}

func ComputerKeyboardWirelessRemoteDeviceComputerWirelessElectronicsQwertyKeyboardBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="3" rx="1"/><path d="M3.5 8.25h7m-7-2.5h7"/></g>`), g.Group(children),
	)
}

func ComputerLaptopDeviceLaptopElectronicsComputerNotebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 8L.72 10.55a1 1 0 0 0 .05 1a1 1 0 0 0 .85.47h10.76a1 1 0 0 0 .85-.47a1 1 0 0 0 0-1L11.5 8m-8-6a1 1 0 0 0-1 1v5h9V3a1 1 0 0 0-1-1Z"/>`), g.Group(children),
	)
}

func ComputerLogoAndroidAndroidCodeAppsBugdroidProgramming(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 13.5v-8a4 4 0 0 1 8 0v8M3 11h8"/><path d="M.5 10V8a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v2M11 .5L9.28 2.22M3 .5l1.72 1.72"/></g>`), g.Group(children),
	)
}

func ComputerLogoAppleOsSystemApple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.18 7.5a2.49 2.49 0 0 1 1.74-2.37a2.91 2.91 0 0 0-4.22-1a1 1 0 0 1-1 0a3.09 3.09 0 0 0-4.41 1.21a5.13 5.13 0 0 0-.54 3.36a7.25 7.25 0 0 0 1.94 4.21a2.09 2.09 0 0 0 2.7.15h0a1.32 1.32 0 0 1 1.57 0a2.06 2.06 0 0 0 2.66-.06a6.58 6.58 0 0 0 1.7-3a2.48 2.48 0 0 1-2.14-2.5Zm-2-5.5L9.68.5"/>`), g.Group(children),
	)
}

func ComputerLogoFacebookOneMediaFacebookSocial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 12.5v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-3V8.76h.71a.61.61 0 0 0 .61-.61v-.77a.61.61 0 0 0-.61-.61h-.67v-.94c0-.84.38-.84.76-.84h.49a.55.55 0 0 0 .43-.18a.58.58 0 0 0 .18-.43v-.74a.62.62 0 0 0-.6-.64H9.65a2.32 2.32 0 0 0-2.39 2.6v1.17h-.64a.61.61 0 0 0-.62.61v.77a.61.61 0 0 0 .62.61h.64v4.74H1.5a1 1 0 0 1-1-1Z"/>`), g.Group(children),
	)
}

func ComputerLogoFacebookTwoMediaFacebookSocial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 4c0-1 .5-1 1-1h.6a.75.75 0 0 0 .76-.76V1.3a.76.76 0 0 0-.77-.76L8.17.52a2.86 2.86 0 0 0-2.95 3.2v1.45h-.79a.76.76 0 0 0-.76.76v.94a.76.76 0 0 0 .76.76h.79V13a.5.5 0 0 0 .5.5h1.81A.5.5 0 0 0 8 13V7.63h.88a.76.76 0 0 0 .76-.76v-.94a.76.76 0 0 0-.76-.76H8Z"/>`), g.Group(children),
	)
}

func ComputerLogoGoogleMediaGoogleSocial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.36 5.83H8.87a.51.51 0 0 0-.51.52v1.41a.51.51 0 0 0 .51.51h1.29a2.75 2.75 0 0 1-3 2.79c-2.24 0-3.32-1.9-3.32-4.06S5 2.94 7.16 2.94a4.07 4.07 0 0 1 2.64.86a.49.49 0 0 0 .72-.22l.63-1.44a.51.51 0 0 0-.15-.63a7.07 7.07 0 0 0-3.8-1C3.56.5 1.08 3.33 1.08 7s2.49 6.5 6.08 6.5s5.76-2.56 5.76-6c0-1.1-.44-1.67-1.56-1.67Z"/>`), g.Group(children),
	)
}

func ComputerLogoLinkedinNetworkLinkedinProfessional(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.54 1.71a1.33 1.33 0 0 1-1.3 1.34A1.36 1.36 0 0 1 .93 1.71A1.34 1.34 0 0 1 2.24.43a1.33 1.33 0 0 1 1.3 1.28ZM1.07 5.43c0-.77.49-.65 1.17-.65s1.16-.12 1.16.65v7.5c0 .78-.49.62-1.16.62s-1.17.16-1.17-.62Zm4.35 0c0-.43.16-.59.41-.64s1.11 0 1.41 0s.42.49.41.86a2.51 2.51 0 0 1 2.24-1a3 3 0 0 1 3.18 3.13v5.12c0 .78-.48.62-1.16.62s-1.16.16-1.16-.62v-4a1.44 1.44 0 0 0-1.52-1.56a1.45 1.45 0 0 0-1.48 1.59v4c0 .78-.49.62-1.17.62s-1.16.16-1.16-.62Z"/>`), g.Group(children),
	)
}

func ComputerLogoPaypalPaymentPaypal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.62 1.59L3.6.51a.49.49 0 0 0-.6.37L1.64 7.29L.63 12a.49.49 0 0 0 .37.58l1.2.26a.49.49 0 0 0 .58-.38l1-4.68l3.44.74a3.52 3.52 0 0 0 4.26-3.43a3.59 3.59 0 0 0-2.86-3.5ZM7.26 6.25l-3-.65l.55-2.6l3 .65a1.32 1.32 0 0 1-.56 2.58Z"/><path d="m5.16 13.5l.62-2.9l3 .63c2.24.49 4.6-1.86 4.63-4.29a3.62 3.62 0 0 0-.2-1.19"/></g>`), g.Group(children),
	)
}

func ComputerLogoSkypeVideoMeetingSkype(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.87 8.63l-.09-.07A6.1 6.1 0 0 0 13 7a6 6 0 0 0-6-6a6.1 6.1 0 0 0-1.55.21l-.07-.09a2.78 2.78 0 0 0-3.89.35a2.78 2.78 0 0 0-.35 3.89l.08.06A5.91 5.91 0 0 0 1 7a6 6 0 0 0 6 6a5.91 5.91 0 0 0 1.58-.22l.06.08a2.78 2.78 0 0 0 3.89-.35a2.78 2.78 0 0 0 .34-3.88Z"/><path fill="currentColor" d="M7.15 4.07c.42 0 1.45.15 1.45.74a.51.51 0 0 1-.49.54c-.31 0-.54-.22-1-.22s-.6.17-.6.48C6.52 6.38 9 5.89 9 7.8a1.76 1.76 0 0 1-1.9 1.71c-.57 0-1.79-.13-1.79-.83a.49.49 0 0 1 .49-.52c.35 0 .76.29 1.25.29a.66.66 0 0 0 .75-.64c0-.87-2.47-.35-2.47-2.06a1.71 1.71 0 0 1 1.82-1.68m0-1a2.7 2.7 0 0 0-2.83 2.68A2.09 2.09 0 0 0 5 7.38a1.52 1.52 0 0 0-.7 1.3c0 1.13 1.07 1.83 2.79 1.83A2.75 2.75 0 0 0 10 7.8a2.33 2.33 0 0 0-.83-1.88a1.59 1.59 0 0 0 .43-1.11c0-1.2-1.27-1.74-2.45-1.74Z"/>`), g.Group(children),
	)
}

func ComputerLogoTwitterMediaTwitterSocial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 9.4a6.77 6.77 0 0 1-2.41 1.21a.5.5 0 0 0 0 .94C8.51 14.39 12.91 10 12.24 5.13l1.12-2.32h-1.3C10.44.82 6.14.92 6.85 5.16c0 0-2.3.41-5.24-2.48A.5.5 0 0 0 .76 3A5.52 5.52 0 0 0 4 9.4Z"/>`), g.Group(children),
	)
}

func ComputerLogoWindowsOneOsSystemMicrosoft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.94 2l12-1.38a.49.49 0 0 1 .56.49v11.82a.5.5 0 0 1-.58.49l-12-1.84a.51.51 0 0 1-.42-.5V2.46A.5.5 0 0 1 .94 2ZM6 1.38v10.98m7.5-5.6H.5"/>`), g.Group(children),
	)
}

func ComputerLogoWindowsTwoOsSystemMicrosoft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 5.82h6.5v-4.5a.49.49 0 0 0-.56-.49L7 1.54Zm-2 0V1.71l-4.06.5a.5.5 0 0 0-.44.5v3.11Zm0 2H.5v3a.51.51 0 0 0 .42.5L5 12Zm2 0v4.4l5.92.95a.5.5 0 0 0 .58-.49V7.82Z"/>`), g.Group(children),
	)
}

func ComputerLogoYoutubeYoutubeClipSocialVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10" x=".5" y="2" rx="2"/><path d="M5.31 9.32v-4.2a.39.39 0 0 1 .6-.34l3.6 2.1a.4.4 0 0 1 0 .69l-3.6 2.1a.4.4 0 0 1-.6-.35Z"/></g>`), g.Group(children),
	)
}

func ComputerMonitorScreenDesktopMonitorDeviceElectronicsDisplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11 5.5h1.5a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H9"/><rect width="8" height="6" x=".5" y="5.5" rx="1"/><path d="M4.5 11.5v2m-2 0h4"/></g>`), g.Group(children),
	)
}

func ComputerMouseComputerDeviceElectronicsMouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="10.5" x="1" y="3" rx="3.5"/><path d="M1 7h7M4.5 7V2.75A2.25 2.25 0 0 1 6.75.5h0A2.25 2.25 0 0 1 9 2.75V3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2V1.5"/></g>`), g.Group(children),
	)
}

func ComputerMouseWirelessRemoteWirelessDeviceElectronicsMouseComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="13" x="2.5" y=".5" rx="4.5"/><path d="M7 3.5V5"/></g>`), g.Group(children),
	)
}

func ComputerPrinterScanDeviceElectronicsPrinterPrintComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 11h2a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h2"/><path d="M3.5 9.5V13a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5V9.5Zm7-5V1a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v3.5M11 7h-1"/></g>`), g.Group(children),
	)
}

func ComputerRobotCyborgArtificialRoboticsRobotIntelligenceMachineTechnologyAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 9.5c0-3-2.46-5-5.5-5s-5.5 2-5.5 5v4h11ZM7 4.5v-4m-5.5 10h11m-7.5 0v3m4-3v3"/><circle cx="5" cy="8" r=".5"/><circle cx="9" cy="8" r=".5"/></g>`), g.Group(children),
	)
}

func ComputerScreenCurveScreenCurvedDeviceElectronicsMonitorDiplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.93 11.34a42.07 42.07 0 0 0-11.86 0a.5.5 0 0 1-.57-.49V2.49A.49.49 0 0 1 1.07 2a42.83 42.83 0 0 0 11.86 0a.49.49 0 0 1 .57.48v8.36a.5.5 0 0 1-.57.5ZM7 10.92v2.5m-2.5 0h5"/>`), g.Group(children),
	)
}

func ComputerScreenOneScreenDeviceElectronicsMonitorDiplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6"/></g>`), g.Group(children),
	)
}

func ComputerScreenTvMoviesTelevisionCathodeCrtTvRayTubeVintageVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10.5" x=".5" y="3" rx="1"/><rect width="8" height="5.5" x="3" y="5.5" rx="1"/><path d="M5 .5L7 3L9 .5"/></g>`), g.Group(children),
	)
}

func ComputerScreenTwoScreenDeviceElectronicsMonitorDiplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="1.5" rx=".5"/><path d="M3.53 13.5a3.51 3.51 0 0 1 6.94 0"/></g>`), g.Group(children),
	)
}

func ComputerSmartWatchOneDeviceTimepieceCirleComputerElectronicsFaceBlankWatchSmart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="3.5"/><path d="M9 4.13V1.5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2.63m4 5.74v2.63a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9.87"/></g>`), g.Group(children),
	)
}

func ComputerSmartWatchTwoDeviceSquareTimepieceComputerElectronicsFaceBlankWatchSmart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9 3.5v-2a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2m4 7v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-2"/><rect width="7" height="7" x="3.5" y="3.5" rx="1"/></g>`), g.Group(children),
	)
}

func ComputerStorageFloppyDiskDiskFloppyElectronicsDeviceDiscComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 12.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V4.91a1 1 0 0 1 .29-.7L4.21.79a1 1 0 0 1 .7-.29h7.59a1 1 0 0 1 1 1Z"/><path d="M10.5 13.5V9a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v4.5m7-13V4a.5.5 0 0 1-.5.5H6a.5.5 0 0 1-.5-.5V.5"/></g>`), g.Group(children),
	)
}

func ComputerStorageHardDiskDeviceDiscDriveComputerDiskElectronicsPlatterTurntableRaid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="m4.5 10.5l1.75-2.75M4.09 7A2.93 2.93 0 0 1 4 5.16a3 3 0 1 1 4.67 3.27M7.5 11H10"/></g>`), g.Group(children),
	)
}

func ComputerStorageHardDriveOneDiskComputerDeviceElectronicsDiscDriveRaid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="7.5" rx="1"/><path d="m.59 8.08l1.74-6.8A1 1 0 0 1 3.3.5h7.4a1 1 0 0 1 1 .78l1.74 6.8M3.5 10.5h3"/><circle cx="10.5" cy="10.5" r=".5"/></g>`), g.Group(children),
	)
}

func ComputerStorageHardDriveTwoDiskComputerDeviceElectronicsDiscDriveRaid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 13.5a3 3 0 0 0 0-6h-7a3 3 0 0 0 0 6Z"/><path d="M.58 9.79L2.17 2.1a2 2 0 0 1 2-1.6h5.7a2 2 0 0 1 2 1.6l1.59 7.69m-9.96.71h3"/><circle cx="10.5" cy="10.5" r=".5"/></g>`), g.Group(children),
	)
}

func ComputerVirtualRealityGamingVirtualGearControllerRealityGamesHeadsetTechnologyVrEyewear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.5H.5v5a1 1 0 0 0 1 1h3a1 1 0 0 0 .78-.38l1.31-1.63a.5.5 0 0 1 .78 0l1.33 1.63a1 1 0 0 0 .78.38h3a1 1 0 0 0 1-1Zm0 0L10.79 2a1 1 0 0 0-.86-.49H4.07a1 1 0 0 0-.86.49L.5 6.5"/>`), g.Group(children),
	)
}

func ComputerVoiceMailMicAudioMikeMusicMicrophone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 6.5a2.5 2.5 0 0 1-5 0V3a2.5 2.5 0 0 1 5 0Z"/><path d="M12 7h0a4.49 4.49 0 0 1-4.5 4.5h-1A4.49 4.49 0 0 1 2 7h0m5 4.5v2"/></g>`), g.Group(children),
	)
}

func ComputerVoiceMailOffMicAudioMikeMusicMicrophoneMuteOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13m-8 8a2.49 2.49 0 0 1-1-2V3a2.5 2.5 0 0 1 5 0v1.5m-.23 3.04A2.5 2.5 0 0 1 8 8.79m-4.42 1.63A4.46 4.46 0 0 1 2 7h0m10 0h0a4.49 4.49 0 0 1-4.5 4.5h-1a5.25 5.25 0 0 1-.56 0m1.06 0v2"/>`), g.Group(children),
	)
}

func ComputerWebcamVideoWorkVideoMeetingCameraCompanyConferenceOffice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.82 3.75L10 5V3.5a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1H9a1 1 0 0 0 1-1V9l2.82 1.25a.5.5 0 0 0 .68-.47V4.22a.5.5 0 0 0-.68-.47Z"/>`), g.Group(children),
	)
}

func ComputerWebcamWebcamCameraFutureTechChatSkypeTechnologyVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.52" r="6"/><path d="M10.07 11.68a7.4 7.4 0 0 1 2.48 1.8m-11.1 0a7.4 7.4 0 0 1 2.48-1.8"/><circle cx="7" cy="6.52" r="3"/><circle cx="7" cy="6.52" r=".5"/></g>`), g.Group(children),
	)
}

func EcologyScienceAlienExtraterristerialLifeFormSpaceUniverseHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 5.5c0 3.59-2.95 8-6.5 8S.5 9.09.5 5.5S3.41.5 7 .5s6.5 1.41 6.5 5Z"/><path d="M2.75 4.75a3.12 3.12 0 0 0 .49 2.44a3.12 3.12 0 0 0 2.44.49a3.12 3.12 0 0 0-.49-2.44a3.12 3.12 0 0 0-2.44-.49Zm8.5 0a3.12 3.12 0 0 1-.49 2.44a3.12 3.12 0 0 1-2.44.49a3.12 3.12 0 0 1 .49-2.44a3.12 3.12 0 0 1 2.44-.49Z"/></g>`), g.Group(children),
	)
}

func EcologyScienceDnaBiologyExperimentLabScience(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.1 3.9a4.65 4.65 0 0 0 3.4.1M4 13.5a4.68 4.68 0 0 0-.1-3.41a4.74 4.74 0 0 1 1-5.16a4.8 4.8 0 0 1 2.46-1.3"/><path d="M6.6 10.37a4.81 4.81 0 0 0 2.47-1.31a4.73 4.73 0 0 0 1-5.16A4.65 4.65 0 0 1 10 .5M.5 10a4.6 4.6 0 0 1 3.4.11m1.04-5.18l4.13 4.13"/></g>`), g.Group(children),
	)
}

func EcologyScienceErlenmeyerFlaskExperimentLabFlaskScienceChemistrySolution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 .5v6l3.59 4.57a1.5 1.5 0 0 1-1.18 2.43H2.59a1.5 1.5 0 0 1-1.18-2.43L5 6.5v-6M3.5.5h7"/>`), g.Group(children),
	)
}

func EcologySciencePlanetSolarSystemRingPlanetSaturnSpaceAstronomy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.63 8.13C.85 10.49.05 12.53.76 13.24c1 1 4.6-1 8-4.44s5.44-7 4.44-8c-.64-.65-2.38 0-4.47 1.41"/><path d="M12.05 4.92A5 5 0 0 1 7.5 12a5.06 5.06 0 0 1-1.95-.39M3.5 10a5 5 0 0 1 7-7"/></g>`), g.Group(children),
	)
}

func EcologyScienceTestTubeExperimentLabScienceChemistryTestTubeSolution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h9M10 .5v10a3 3 0 0 1-6 0V.5"/>`), g.Group(children),
	)
}

func EntertainmentCameraVideoFilmTelevisionTvCameraMoviesVideoRecorder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.5" cy="3.5" r="3"/><circle cx="10.5" cy="4.5" r="2"/><rect width="7.5" height="4.5" x="3.5" y="9" rx="1"/><path d="M13.5 10v2.5"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonFastForwardOneButtonControlsFastForwardMoviesTelevisionVideoTv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L1.6 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.62Z"/><path d="M7 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L8.1 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.6Z"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonFastForwardTwoButtonControlsFastForwardMoviesTelevisionVideoTv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L1.6 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.62Z"/><path d="M7 10.29a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L8.1 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.6Zm6.5-7.79v9"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonMoveCircleMoveButtonCircleDirectionArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M5.5 4L7 3l1.5 1m-3 6L7 11l1.5-1M10 5.5L11 7l-1 1.5m-6-3L3 7l1 1.5"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonNextButtonTelevisionButtonsMoviesSkipNextVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5v13m-13-1.84a1 1 0 0 0 .52.88a1 1 0 0 0 1 0l7.19-4.7a1 1 0 0 0 0-1.68L2 1.5a1 1 0 0 0-1 0a1 1 0 0 0-.52.88Z"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonNextCircleCircleMultiMultimediaButtonSkipMediaNextControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M9.5 5v4m-5-4.5l3 2.5l-3 2.5v-5z"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonPauseCircleControlsPauseMultiMediaMultimediaButtonCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M5.5 4.5v5m3-5v5"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonPauseOneButtonTelevisionButtonsMoviesTvPauseVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75.52v13m4.5-13v13"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonPauseTwoButtonTelevisionButtonsMoviesTvPauseVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4.5" height="13" x=".5" y=".5" rx="1"/><rect width="4.5" height="13" x="9" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonPlayButtonTelevisionButtonsMoviesPlayTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 12.35a1.14 1.14 0 0 0 .63 1a1.24 1.24 0 0 0 1.22 0L12 8a1.11 1.11 0 0 0 0-2L3.35.69a1.24 1.24 0 0 0-1.22 0a1.14 1.14 0 0 0-.63 1Z"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonPlayCircleControlsMediaMultiPlayMultimediaButtonCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m5.5 4.5l4 2.5l-4 2.5v-5z"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonPlayPauseButtonTelevisionButtonsMoviesPlayPauseVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2.5v9m-3-9v9m-10-1.21a.7.7 0 0 0 .37.62a.71.71 0 0 0 .73 0l5.08-3.3a.7.7 0 0 0 0-1.18L1.6 3.11a.71.71 0 0 0-.73 0a.7.7 0 0 0-.37.62Z"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonPowerOnePowerButtonOnOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5V6m4-4.12a6.5 6.5 0 1 1-8 0"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonPowerTwoPowerButtonOnOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 4.5v5"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonPreviousButtonTelevisionButtonsMoviesSkipPreviousVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5v13m13-1.84a1 1 0 0 1-.52.88a1 1 0 0 1-1 0l-7.19-4.7a1 1 0 0 1 0-1.68L12 1.5a1 1 0 0 1 1 0a1 1 0 0 1 .52.88Z"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonPreviousCirclePreviousControlsMultiMediaMultimediaButtonCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M4.5 5v4m5-4.5L6.5 7l3 2.5v-5z"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonRecordOneButtonMultimediaMultiMediaControlsRecordCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r="2.5"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonRecordThreeButtonTelevisionButtonsMoviesRecordTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="7" cy="7" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonRewindOneRewindTelevisionButtonMoviesButtonsTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L7.32 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.62Z"/><path d="M7 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L.82 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.6Z"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonRewindTwoRewindTelevisionButtonMoviesButtonsTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L7.32 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.62Z"/><path d="M7 10.29a.7.7 0 0 1-.37.62a.71.71 0 0 1-.73 0L.82 7.59a.7.7 0 0 1 0-1.18l5.08-3.3a.71.71 0 0 1 .73 0a.7.7 0 0 1 .37.6ZM.5 2.5v9"/></g>`), g.Group(children),
	)
}

func EntertainmentControlButtonStopButtonTelevisionButtonsMoviesStopTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="13" height="13" x=".5" y=".52" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`), g.Group(children),
	)
}

func EntertainmentControlButtonStopCircleStopControlsMultiMediaMultimediaButtonCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><rect width="5" height="5" x="4.5" y="4.5" rx="1"/></g>`), g.Group(children),
	)
}

func EntertainmentEarpodsAirpodsAudioEarpodsMusicEarbudsTrueWireless(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5a3 3 0 0 0-2 .76v4.48a3 3 0 0 0 2 .76a2.74 2.74 0 0 0 .5 0v5.79a1.25 1.25 0 0 0 2.5 0V3.5a3 3 0 0 0-3-3Zm9 0a3 3 0 0 1 2 .76v4.48a3 3 0 0 1-2 .76a2.74 2.74 0 0 1-.5 0v5.79a1.25 1.25 0 0 1-2.5 0V3.5a3 3 0 0 1 3-3Z"/>`), g.Group(children),
	)
}

func EntertainmentMusicNoteOneMusicAudioNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.25" cy="11" r="2.5"/><path d="M6.75 11V.5h0a5.5 5.5 0 0 1 5.5 5.5h0"/></g>`), g.Group(children),
	)
}

func EntertainmentMusicNoteTwoMusicAudioNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="11.42" r="2"/><circle cx="11.5" cy="8.92" r="2"/><path d="M13.5 8.92V1.08a.5.5 0 0 0-.63-.48l-8 2.22a.5.5 0 0 0-.37.48v8.12m0-5.5l9-2.5"/></g>`), g.Group(children),
	)
}

func EntertainmentNewsPaperNewspaperPeriodicalFoldContent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 4.5V11a1.25 1.25 0 0 1-1.25 1.25h0A1.25 1.25 0 0 1 11 11h0V2.25a.5.5 0 0 0-.5-.5H1a.5.5 0 0 0-.5.5v9a1 1 0 0 0 1 1h10.75"/><path d="M3.5 4.25H8v2.5H3.5zm0 5.5H8"/></g>`), g.Group(children),
	)
}

func EntertainmentPartyPopperHobbyEntertainmentPartyPopperConfettiEvent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.85 13.2l-6.68-2.49a1.25 1.25 0 0 1-.48-2.05l4.19-4.19a1.26 1.26 0 0 1 2.06.53l2.48 6.68a1.22 1.22 0 0 1-1.57 1.52Zm-9.8-6.07a2.06 2.06 0 0 1 1.46-.21m.82-2.64A2.1 2.1 0 0 1 4 2.83M6.63.72A4.72 4.72 0 0 0 6.76 4"/><circle cx="1" cy="3.28" r=".5"/></g>`), g.Group(children),
	)
}

func EntertainmentPlayListFourScreenTelevisionDisplayPlayerMoviesPlayersTvMediaVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5.23 9.23V4.91a.41.41 0 0 1 .62-.35l3.7 2.15a.41.41 0 0 1 0 .71l-3.7 2.16a.41.41 0 0 1-.62-.35Z"/></g>`), g.Group(children),
	)
}

func EntertainmentPlayListOneScreenTelevisionDisplayPlayerMoviesMovieTvMediaPlayersVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5.49 10.56V6.73A.36.36 0 0 1 6 6.42l3.32 1.91a.37.37 0 0 1 0 .63L6 10.88a.37.37 0 0 1-.51-.32ZM.5 4h13M4 4L5.5.5m3 3.5L10 .5"/></g>`), g.Group(children),
	)
}

func EntertainmentPlayListThreePlayerTelevisionDisplayMoviesSmartphoneMediaTvPlayersScreenVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8.5" height="13" x="2.75" y=".5" rx="1"/><path d="M5.27 7.14V3.82a.31.31 0 0 1 .47-.28l2.85 1.67a.31.31 0 0 1 0 .54L5.74 7.41a.31.31 0 0 1-.47-.27ZM2.75 10.5h8.5"/></g>`), g.Group(children),
	)
}

func EntertainmentPlayListTwoPlayerTelevisionMoviesSliderMediaTvPlayersVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 8.68V5.32c0-.25.23-.4.41-.28l2.45 1.68a.35.35 0 0 1 0 .56L5.91 9c-.18.08-.41-.07-.41-.32Z"/><rect width="8" height="11" x="3" y="1.5" rx="1"/><path d="M.5 3v8m13-8v8"/></g>`), g.Group(children),
	)
}

func EntertainmentRadioAntennaAudioMusicRadio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9.5" x=".5" y="4" rx="1"/><path d="M2 4L13.5.5"/><circle cx="4.5" cy="10" r="1.5"/><path d="M9.5 8.75H11m-1.5 2.5H11M.5 6.5h13"/></g>`), g.Group(children),
	)
}

func EntertainmentRecordingTapeOneFilmTelevisionTvMoviesReelVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.5" cy="5.5" r="5"/><circle cx="5.5" cy="5.5" r="1.5"/><path d="M10.5 5.5V12a1.5 1.5 0 0 0 1.5 1.5h0a1.5 1.5 0 0 0 1.5-1.5v-.5"/></g>`), g.Group(children),
	)
}

func EntertainmentRecordingTapeTwoPhoneDeviceMailMobileFaxVoiceMachineAnswering(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3" cy="7" r="2.5"/><circle cx="11" cy="7" r="2.5"/><path d="M3 9.5h8"/></g>`), g.Group(children),
	)
}

func EntertainmentSpeakerOneSpeakerMusicAudioSubwoofer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.43 3.43a1.5 1.5 0 1 0-1.86-1.86a6.49 6.49 0 0 0-7.14 0a1.5 1.5 0 1 0-1.86 1.86A6.52 6.52 0 0 0 .5 7a6.43 6.43 0 0 0 .83 3.17A1.49 1.49 0 0 0 2 13a1.51 1.51 0 0 0 1.26-.69a6.47 6.47 0 0 0 7.48 0a1.5 1.5 0 0 0 2.76-.81a1.49 1.49 0 0 0-.83-1.33A6.43 6.43 0 0 0 13.5 7a6.52 6.52 0 0 0-1.07-3.57Z"/><circle cx="7" cy="7" r="3.5"/><circle cx="7" cy="7" r=".5"/></g>`), g.Group(children),
	)
}

func EntertainmentSpeakerTwoSpeakersMusicAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="13" x="2" y=".5" rx="1"/><circle cx="7" cy="9" r="2.5"/><circle cx="7" cy="3.5" r=".5"/><circle cx="7" cy="9" r=".5"/></g>`), g.Group(children),
	)
}

func EntertainmentTicketHobbyTicketEventEntertainmentStubTheater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.65 2.24l-.39.09a1.19 1.19 0 0 1-.91 1.43a1.22 1.22 0 0 1-1.44-.92L1.13 3a.81.81 0 0 0-.61 1s.36 1.67.78 3.55m10.85-5.32a1.32 1.32 0 1 1-2.44-1L8.9.85a.89.89 0 0 0-1.16.47L4.38 9.23l-.44 1a.89.89 0 0 0 .47 1.16l4.06 1.73a.89.89 0 0 0 1.16-.47l3.8-8.94A.89.89 0 0 0 13 2.58Z"/>`), g.Group(children),
	)
}

func EntertainmentVolumeDownSpeakerDownVolumeControlAudioMusicDecrease(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5m7.5 1.5h3"/>`), g.Group(children),
	)
}

func EntertainmentVolumeLevelHighSpeakerHighVolumeControlAudioMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5m9.5-1a4.38 4.38 0 0 1 1 3a6.92 6.92 0 0 1-1 3.5m-2-5A2.19 2.19 0 0 1 11 7a2.19 2.19 0 0 1-.5 1.5"/>`), g.Group(children),
	)
}

func EntertainmentVolumeLevelLowVolumeSpeakerLowerDownControlMusicLowAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5m7.5.5A2.19 2.19 0 0 1 11 7a2.19 2.19 0 0 1-.5 1.5"/>`), g.Group(children),
	)
}

func EntertainmentVolumeLevelOffVolumeSpeakerControlMusicAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5"/>`), g.Group(children),
	)
}

func EntertainmentVolumeMuteSpeakerRemoveVolumeControlAudioMusicMuteOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5m10.5.5l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func EntertainmentVolumeOffSpeakerMusicMuteVolumeControlAudioOffMute(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13M4.5 5H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1.5ZM10 4V3a1 1 0 0 0-.55-.89a1 1 0 0 0-1 .08L4.5 5m2.17 5.56l1.74 1.25a1 1 0 0 0 1 .08A1 1 0 0 0 10 11V7M4.5 9l.29.21"/>`), g.Group(children),
	)
}

func EntertainmentVolumeUpVolumeSpeakerUpControlMusicPlusAddAudioIncrease(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 5v3m-1.5-1.5h3M3 5H1.5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1H3Zm0 4l3.91 2.81a1 1 0 0 0 1 .08A1 1 0 0 0 8.5 11V3a1 1 0 0 0-.5-.89a1 1 0 0 0-1 .08L3 5"/>`), g.Group(children),
	)
}

func EntertainmentWalkManPlayerTapesTapeHeadphonesMusicWalkmanHeadsetAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="7" x="4.5" y="6.5" rx="1"/><path d="M2 7.5v3m10-3v3m-10-2H1A.5.5 0 0 1 .5 8V7A6.5 6.5 0 0 1 7 .5h0A6.5 6.5 0 0 1 13.5 7v1a.5.5 0 0 1-.5.5h-1M7 9v1"/></g>`), g.Group(children),
	)
}

func FoodBaguetteCookBreadGlutenDrinkCookingNutritionBaguetteFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.636 6.93l6.468-5.074a3 3 0 0 1 4.212.509l.975 1.243a1 1 0 0 1-.17 1.404l-9.615 7.542a1 1 0 0 1-1.404-.17l-.975-1.243a3 3 0 0 1 .51-4.212Zm1.234-.97l1.56 2m2.03-4.82l1.57 2"/>`), g.Group(children),
	)
}

func FoodBarbequePotCookGrillBbqDrinkCookingNutritionPotBarbecueGrillingFoodCauldron(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 12a6.51 6.51 0 0 0 6.48-6a.5.5 0 0 0-.48-.5H1a.5.5 0 0 0-.48.5A6.51 6.51 0 0 0 7 12Zm-2.95-.71L3.5 13.5m6.45-2.21l.55 2.21M3.5 3V2m7 1V2M7 3V.5"/>`), g.Group(children),
	)
}

func FoodBurgerDrinkBurgerFastCookCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 .5h6A3.5 3.5 0 0 1 13.5 4v0a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v0A3.5 3.5 0 0 1 4 .5Zm-3.5 7h13M13 10H7l-1.5 1.5l-3-1.5H1a.5.5 0 0 0-.5.5h0a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3h0a.5.5 0 0 0-.5-.5Z"/>`), g.Group(children),
	)
}

func FoodCakeCandleBirthdayEventSpecialSweetCakeBake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="1"/><path d="M3 3.5v-2m4 2v-2m4 2v-2m2.5 8.5H13a2 2 0 0 1-2-2a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-2 2H.5"/></g>`), g.Group(children),
	)
}

func FoodCandyCaneCandySweetCaneChristmas(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5a1.5 1.5 0 0 1 3 0v7a1.5 1.5 0 0 0 3 0V5a4.5 4.5 0 0 0-9 0v1a1.5 1.5 0 0 0 3 0Zm3 5l3-3M8.42 4.51L9.95 1.6M5.86 4.02L3.03 2.88"/>`), g.Group(children),
	)
}

func FoodCheeseOneCookCheeseAnimalProductsCookingNutritionDairyFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.59 5.75H13M9.09.77L9 .75L.75 5.6A.5.5 0 0 0 .5 6v1.75c1 0 2.5 0 2.5 1.75S1.47 11.25.5 11.25v1.5a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-7A5.88 5.88 0 0 0 9.09.77Z"/><circle cx="10" cy="8.25" r=".5"/><circle cx="6.5" cy="10.25" r=".5"/></g>`), g.Group(children),
	)
}

func FoodCheeseTwoCookCheeseAnimalProductsCookingNutritionDairyFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.09.77L9 .75L.75 5.61A.47.47 0 0 0 .5 6v6.71a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-7A5.88 5.88 0 0 0 9.09.77ZM.59 5.75H13"/><path d="M.5 10.42h0l.08-.07A2 2 0 0 1 2 9.75a2 2 0 0 1 2 2a2 2 0 0 1-.59 1.41l-.08.09"/><circle cx="9" cy="9.25" r="1.5"/></g>`), g.Group(children),
	)
}

func FoodCherriesCookPlantCherryPlantsDrinkCookingNutritionVegetarianFruitFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.36" cy="10.9" r="2.6"/><circle cx="10.64" cy="9.86" r="2.6"/><path d="M9.41 7.57A10.36 10.36 0 0 1 6 .5a13.78 13.78 0 0 0-2.6 7.8M3.36.5h5.2"/></g>`), g.Group(children),
	)
}

func FoodDrinksBeerMugBeerCookBreweryDrinkMugCookingNutritionBrewBrewingFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5.5h8a1 1 0 0 1 1 1v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a1 1 0 0 1 1-1Zm9 2.5h2a1 1 0 0 1 1 1v4.5a1 1 0 0 1-1 1h-2M3.75 4v5.5M7.25 4v5.5"/>`), g.Group(children),
	)
}

func FoodDrinksCocktailGlassCookAlcoholFoodCocktailDrinkCookingAlcoholicBeverageGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.71.5a.48.48 0 0 0-.46.31a.49.49 0 0 0 .1.54L7 7l5.65-5.65a.49.49 0 0 0 .1-.54a.48.48 0 0 0-.46-.31ZM7 7v6.5m-3 0h6"/>`), g.Group(children),
	)
}

func FoodDrinksCocktailOneCookAlcoholFoodCocktailDrinkCookingNutritionAlcoholicBeverageGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 8.5L2.48 1.26A.5.5 0 0 1 2.9.5h8.2a.5.5 0 0 1 .42.76Zm0 0v5m-2.5 0h5M3.56 3h6.88"/>`), g.Group(children),
	)
}

func FoodDrinksCocktailShakerAlcoholDrinkMixShakeCocktail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 6l-.89 6.62a1 1 0 0 1-1 .88H4.38a1 1 0 0 1-1-.88L2.5 6m8.23-2.32a1 1 0 0 0-1-.68H4.22a1 1 0 0 0-.95.68L2.5 6h9Z"/><path d="M8.5 3V1.5a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1V3"/></g>`), g.Group(children),
	)
}

func FoodDrinksCoffeeMugCoffeeCookCupDrinkMugCookingNutritionCafeCaffeineFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5.5h5a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5a1 1 0 0 1 1-1Zm6 1h.5a2.5 2.5 0 0 1 0 5H9M4 .5v2m3-2v2"/>`), g.Group(children),
	)
}

func FoodDrinksMilkCanisterMilkDairyCanisterDrink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11 7.5l-1.5-2h-5L3 7.5v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1Zm0 0l2.5-2M3 7.5l-2.5-2m4-5h5a1 1 0 0 1 1 1V3h0h-7h0V1.5a1 1 0 0 1 1-1Zm0 2.5v2.5m5-2.5v2.5"/>`), g.Group(children),
	)
}

func FoodDrinksTeaCupHerbalCookTeaTisaneCupDrinkCookingNutritionMugFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 4h8a.5.5 0 0 1 .5.5V9a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2V4.5A.5.5 0 0 1 2 4ZM.5 13.5h13m-3-8.5h1a2 2 0 0 1 2 2h0a2 2 0 0 1-2 2h-1M3 .5v1m5-1v1M5.5.5v1"/>`), g.Group(children),
	)
}

func FoodDrinksTeapotTeaPotDrinkHotHerb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 6a5.51 5.51 0 0 0-4.46 2.28L1.5 7l-1 1l2 3.5v1a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-1A5.5 5.5 0 0 0 8 6Zm0 0V4.5"/><path d="M4 7.73V4.5a4 4 0 0 1 8 0v3.23"/></g>`), g.Group(children),
	)
}

func FoodDrinksWaterGlassGlassWaterJuiceDrinkLiquid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.2 11.73a2 2 0 0 1-2 1.77H4.78a2 2 0 0 1-2-1.77L1.5.5h11ZM1.96 4.5h10.08"/>`), g.Group(children),
	)
}

func FoodDrinksWineBottleCookBottleWineDrinkCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.25 4V1a.5.5 0 0 0-.5-.5h-1.5a.5.5 0 0 0-.5.5v3h0a3.36 3.36 0 0 0-1.5 2.8v5.7a1 1 0 0 0 1 1h3.5a1 1 0 0 0 1-1V6.8A3.36 3.36 0 0 0 8.25 4Zm-4 3.5h5.5m0 3.5h-5.5"/>`), g.Group(children),
	)
}

func FoodDrinksWineGlassDrinkCookGlassCookingWineNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.67.5a1 1 0 0 0-1 .89C2.56 2.45 2.5 4.05 2.5 4.5c0 2.62 2 4 4.5 4s4.5-1.38 4.5-4c0-.45-.06-2-.17-3.11a1 1 0 0 0-1-.89ZM7 8.5v5m-2 0h4"/>`), g.Group(children),
	)
}

func FoodDrumStickOneCookAnimalDrumsticksProductsChickenCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.49 9.6a1.66 1.66 0 0 0-3-.79l-1-1s-.08-.05-.11-.08A5.6 5.6 0 0 0 8 1.87a4.15 4.15 0 0 0-6.13 0A4.16 4.16 0 0 0 1.83 8a5.58 5.58 0 0 0 6 1.34a.56.56 0 0 0 0 .08L9 10.58A1.63 1.63 0 0 0 8.43 12a1.66 1.66 0 0 0 3.31-.34a1.6 1.6 0 0 0-.05-.21H12a1.65 1.65 0 0 0 1.49-1.85Z"/>`), g.Group(children),
	)
}

func FoodFishCookCookingFishSeafoodNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.23 11.06a5 5 0 0 0 2.63-2.14a2.88 2.88 0 0 0 1.25 1A1 1 0 0 0 13.5 9V6.39a1 1 0 0 0-1.39-.92a2.83 2.83 0 0 0-1.25 1c-1-1.56-3.72-2.77-6.45-2.77a4 4 0 0 0-3.91 4a4 4 0 0 0 4.16 3.75h1.08"/><circle cx="4" cy="6.69" r=".5"/><path d="M3.23 3.87C4 2.68 5.81 1.34 9.82 1.8a5.16 5.16 0 0 0 0 3.57M4.64 9.53a3.5 3.5 0 0 0 4 2.76L7.63 9"/></g>`), g.Group(children),
	)
}

func FoodIceCreamOneCookFrozenBitePopsicleCreamIceCookingNutritionFreezerColdFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 10v2.5a1 1 0 0 1-1 1h0a1 1 0 0 1-1-1V10m3.75-7A1.75 1.75 0 0 1 8 1.25a1.74 1.74 0 0 1 .1-.56A3.63 3.63 0 0 0 7 .5A3.5 3.5 0 0 0 3.5 4v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V4a3.63 3.63 0 0 0-.19-1.1a1.74 1.74 0 0 1-.56.1Z"/>`), g.Group(children),
	)
}

func FoodIceCreamThreeCookFrozenConeCreamIceCookingNutritionFreezerColdFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.36 6.48l-2.87 6.7a.54.54 0 0 1-1 0l-2.85-6.7"/><circle cx="4.88" cy="4.75" r="2.13"/><path d="M4.88 2.63a2.12 2.12 0 1 1 4.24 0"/><circle cx="9.13" cy="4.75" r="2.13"/></g>`), g.Group(children),
	)
}

func FoodIceCreamTwoCookFrozenFoodPopsicleFreezerNutritionCreamStickColdIceCooking(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="9" x="3" y=".5" rx="1"/><path d="M5.5 3v3.5m3-3.5v3.5m0 3V12a1.5 1.5 0 0 1-3 0V9.5"/></g>`), g.Group(children),
	)
}

func FoodKitchenwareBowlChopStickCookSoupBowlChopsticksCookingNutritionAsianFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5A6.5 6.5 0 0 0 13.5 7H.5A6.5 6.5 0 0 0 7 13.5ZM7.5 5l3-4.5M9 5.5l3.5-3"/>`), g.Group(children),
	)
}

func FoodKitchenwareChefToqueHatCookGearChefCookingNutritionToolsClothesHatClothingFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 13.5h9m1.93-10.1a2.49 2.49 0 0 0-4.09-1.26a2.49 2.49 0 0 0-4.68 0A2.49 2.49 0 0 0 .57 3.4A2.51 2.51 0 0 0 2.5 6.45V10a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V6.45a2.51 2.51 0 0 0 1.93-3.05ZM2.5 8.5h9"/>`), g.Group(children),
	)
}

func FoodKitchenwareForkSpoonForkSpoonFoodDineCookUtensilsEatRestaurantDining(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="10.6" cy="3.5" rx="2.4" ry="3"/><path d="M10.6 6.5v7M3.5.5v13M6 .5V3a2.5 2.5 0 0 1-2.5 2.5h0A2.5 2.5 0 0 1 1 3V.5"/></g>`), g.Group(children),
	)
}

func FoodKitchenwareMicrowaveCookFoodAppliancesCookingNutritionApplianceMicrowave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx=".5"/><rect width="6.5" height="5" x="2.5" y="4.5" rx=".5"/><path d="M11 5h.5M11 6.5h.5M11.25 9h0"/></g>`), g.Group(children),
	)
}

func FoodKitchenwareNoFoodAllowedForkSpoonFoodDineCookUtensilsEatRestaurantNotAllowed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5.5v10M6 .5V3a2.5 2.5 0 0 1-2.5 2.5h0A2.5 2.5 0 0 1 1 3V.5m-.5 13l13-13m-1 5A14.61 14.61 0 0 1 13 10H8.5m0-4.5v-5a7.41 7.41 0 0 1 2.66 2.34M8.5 13.5V9"/>`), g.Group(children),
	)
}

func FoodKitchenwareRefrigeratorFridgeCookAppliancesCookingNutritionFreezerApplianceFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="12" x="2.5" y=".5" rx="1"/><path d="M2.5 6h9m-7-3v.5m-.5 9v1m6-1v1"/></g>`), g.Group(children),
	)
}

func FoodKitchenwareServingDomeCookToolDomeKitchenServingPaltterDishToolsFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 3h0a6.5 6.5 0 0 1 6.5 6.5v0a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v0A6.5 6.5 0 0 1 7 3Zm0 0V1.5m-6.5 11h13"/>`), g.Group(children),
	)
}

func FoodKitchenwareSpoonPlateForkPlateFoodDineCookUtensilsEatRestaurantDining(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v2.25a2.5 2.5 0 0 0 5 0V.5M3 .5v13m10.5 0a6.5 6.5 0 0 1 0-13"/><path d="M13.5 10.5a3.5 3.5 0 0 1 0-7"/></g>`), g.Group(children),
	)
}

func FoodMeatCookMeatCownSliceOrganicCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 5.44V4.5a2 2 0 0 1 2-2h6.05a.91.91 0 0 1 .95.84v3.82a.91.91 0 0 1-.95.84H8.8"/><path d="M13.5 9.24v2.88a1.38 1.38 0 0 1-1.38 1.38H6.79a1.38 1.38 0 0 1-1.33-1h0a1.37 1.37 0 0 0-1.32-1h-.56A3.08 3.08 0 0 1 .5 8.42V4.56"/><path d="M11.5.5a2 2 0 0 1 2 2v6.74a1.38 1.38 0 0 1-1.38 1.38H8a1.38 1.38 0 0 1-1.33-1h0a1.37 1.37 0 0 0-1.32-1h-.79a4.06 4.06 0 0 1 0-8.12Z"/><circle cx="5.33" cy="5.14" r=".5"/></g>`), g.Group(children),
	)
}

func FoodPizzaDrinkCookFastCookingNutritionPizzaFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.75.5L8 6l5.5-1.75A3.75 3.75 0 0 0 9.75.5Z"/><path d="M7.5.52L7 .5A6.5 6.5 0 1 0 13.5 7v-.5"/><circle cx="5.5" cy="9.5" r=".5"/><circle cx="4" cy="5" r=".5"/><circle cx="10.5" cy="8" r=".5"/></g>`), g.Group(children),
	)
}

func FoodPopcornCookCornMovieSnackCookingNutritionBake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.56 6l.83 6.62a1 1 0 0 0 1 .88h5.23a1 1 0 0 0 1-.88L11.44 6"/><path d="M11 3.07a1.49 1.49 0 0 0-.36.05a1.81 1.81 0 0 0 .14-.69A2 2 0 0 0 8.76.5A2 2 0 0 0 7 1.56A2 2 0 0 0 5.24.5a2 2 0 0 0-2 1.93a1.81 1.81 0 0 0 .14.69A1.49 1.49 0 0 0 3 3.07a1.5 1.5 0 1 0 0 3a1.56 1.56 0 0 0 1.2-.56a1.53 1.53 0 0 0 1.44 1A1.55 1.55 0 0 0 7 5.76a1.55 1.55 0 0 0 1.32.74a1.53 1.53 0 0 0 1.44-1a1.56 1.56 0 0 0 1.2.56a1.5 1.5 0 1 0 0-3Z"/></g>`), g.Group(children),
	)
}

func FoodSteakCookGrillSteakBbqCookingNutritionBarbecueGrillingFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.62 2.5C4.29 2.5.5 3.15.5 6a2.75 2.75 0 0 0 3 2.44a3.55 3.55 0 0 0 1.28-.24a1 1 0 0 1 .8 0a6.41 6.41 0 0 0 3.04.8c2.7 0 4.88-1.46 4.88-3.25S11.32 2.5 8.62 2.5Z"/><path d="M13.5 5.75v2.5c0 1.79-2.18 3.25-4.88 3.25a6.41 6.41 0 0 1-3.06-.73a1 1 0 0 0-.8 0a3.55 3.55 0 0 1-1.28.23a2.75 2.75 0 0 1-3-2.44V6"/><circle cx="4" cy="5.5" r=".5"/></g>`), g.Group(children),
	)
}

func FoodToastBreadToastBreakfast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 3A2.5 2.5 0 0 0 11 .5H3A2.5 2.5 0 0 0 1.5 5v7.5a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V5a2.49 2.49 0 0 0 1-2Zm-5 .5l-4 4m5-1l-4 4"/>`), g.Group(children),
	)
}

func FoodWaterMelonCookPlantPlantsDrinkCookingNutritionWatermelonFruitVegetarianFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.13 1L1 10.65a1.14 1.14 0 0 0 .48 1.55a12.22 12.22 0 0 0 11 0a1.14 1.14 0 0 0 .52-1.55L7.87 1a1 1 0 0 0-1.74 0Z"/><path d="M12.24 9.29a12 12 0 0 1-10.48 0M5.5 7.5V8m3-.5V8M7 4.5V5"/></g>`), g.Group(children),
	)
}

func FoodWheatCookPlantBreadGlutenGrainCookingNutritionFoodWheat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.53 10.85l.84 2.35M1.98.97l.34.94m3.21 8.94l-1.75-.43m1.75.43L6.6 9.4m-1.92-.9l-1.75-.44m1.75.44l1.08-1.45m-1.92-.91l-1.75-.43m1.75.43L4.91 4.7m-1.92-.91l-1.75-.43m1.75.43l1.08-1.45m5.44 7.91l-.56 2.25M11.85.8l-.22.9m-2.12 8.55L8.38 9.01m1.13 1.24l1.57-.56M10.06 8L8.94 6.76M10.06 8l1.58-.57m-1.02-1.68L9.5 4.51m1.12 1.24l1.58-.57M11.18 3.5l-1.12-1.24m1.12 1.24l1.58-.57"/>`), g.Group(children),
	)
}

func HealthMedicalHeartRateHealthBeautyInformationDataBeatPulseMonitorHeartRateInfo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.58 4.31C1.09 1.85 4.12 0 7 3.27c4.11-4.71 8.5 1.13 5.52 4.14L7 12.5l-3.23-3"/><path d="M.5 7H3l1.5-2l2 3.5l1.5-2h1.5"/></g>`), g.Group(children),
	)
}

func HealthMedicalRibbonOneRibbonMedicalCancerHealthBeautySymbol(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 13.5S10 6.39 10 3.5a2.84 2.84 0 0 0-3-3a2.84 2.84 0 0 0-3 3c0 2.89 6.5 10 6.5 10"/>`), g.Group(children),
	)
}

func HealthMedicalRibbonTwoRibbonMedicalCancerHealthBeautySymbol(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 3.14v1.5m1.51 2.51a3.62 3.62 0 0 0 1.49-3A3.29 3.29 0 0 0 7 .64a3.28 3.28 0 0 0-3 3.5a3.62 3.62 0 0 0 1.49 3l-3.21 3.68a.51.51 0 0 0 0 .65l1.37 1.7a.47.47 0 0 0 .38.19a.52.52 0 0 0 .39-.17A23 23 0 0 0 7 9.64a23 23 0 0 0 2.6 3.55a.52.52 0 0 0 .39.17a.47.47 0 0 0 .38-.19l1.37-1.7a.5.5 0 0 0 0-.64Z"/>`), g.Group(children),
	)
}

func HealthMedicalSyringeInstrumentMedicalSyringeHealthBeautyNeedle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 3L2.42 7.49a2 2 0 0 0 0 2.83l1.26 1.26a2 2 0 0 0 2.83 0L11 7M10.5.5l3 3M9 5l3-3m-8.95 8.95L.5 13.5m5-12l7 7"/>`), g.Group(children),
	)
}

func ImageAccessoriesLensesPhotosCameraShutterPicturePhotographyPicturesPhotoLens(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r="2.5"/><path d="M4.5 7V1M7 4.5h6M9.5 7v6M7 9.5H1"/></g>`), g.Group(children),
	)
}

func ImageCameraOnePhotosPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 5a1 1 0 0 0-1-1h-2L9 2H5L3.5 4h-2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1Z"/><circle cx="7" cy="7.5" r="2.25"/></g>`), g.Group(children),
	)
}

func ImageCameraSettingPinPhotosCameraMapPhotographyPicturesMapsSettingsPinPhoto(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 3.5a1 1 0 0 0-1-1h-2L9 .5H5l-1.5 2h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h3l2.5 4l2.5-4h3a1 1 0 0 0 1-1Z"/><circle cx="7" cy="5.5" r="2"/></g>`), g.Group(children),
	)
}

func ImageCameraTripodTripodPhotosPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 8.5v5m0-5l-4.5 5m4.5-5l4.5 5"/><rect width="11" height="8" x="1.5" y=".5" rx="1"/><circle cx="7" cy="4.5" r="1.5"/></g>`), g.Group(children),
	)
}

func ImageCameraTwoPhotosPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="3.75" rx="1"/><circle cx="7" cy="8.25" r="2"/><path d="M9.5 3.75v-1.5a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v1.5"/></g>`), g.Group(children),
	)
}

func ImageFlashOffFlashPowerConnectChargeOffElectricityLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13M5.19 5.19L8.5.5V5h3L8.81 8.81m-1.66 2.35L5.5 13.5V9h-3l1.41-2"/>`), g.Group(children),
	)
}

func ImageFlashOneFlashPowerConnectChargeElectricityLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.25.5L2 5.81a.5.5 0 0 0 .46.69h2.79l-2 7l8.59-8.14a.5.5 0 0 0-.34-.86H7.75l2-4Z"/>`), g.Group(children),
	)
}

func ImageFlashTwoFlashPowerConnectChargeElectricityLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 .5v5h3.5l-5.5 8v-5H2.5L8 .5z"/>`), g.Group(children),
	)
}

func ImagePhotoCompositionOvalCameraFrameCompositionPhotographyPicturesLandscapePhotoOval(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 11a10.34 10.34 0 0 0 1-4a10.34 10.34 0 0 0-1-4A13.27 13.27 0 0 0 7 2a13.27 13.27 0 0 0-5.5 1a10.34 10.34 0 0 0-1 4a10.34 10.34 0 0 0 1 4A13.27 13.27 0 0 0 7 12a13.27 13.27 0 0 0 5.5-1Z"/>`), g.Group(children),
	)
}

func ImagePhotoCompositionVerticalCameraPortraitFrameVerticalCompositionPhotographyPhoto(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.2 1.19a.52.52 0 0 0 0-.47a.49.49 0 0 0-.46-.22H3.26a.49.49 0 0 0-.41.22a.52.52 0 0 0-.05.47a14.67 14.67 0 0 1 0 11.62a.52.52 0 0 0 .05.47a.49.49 0 0 0 .41.22h7.48a.49.49 0 0 0 .41-.22a.52.52 0 0 0 0-.47a14.67 14.67 0 0 1 .05-11.62Z"/>`), g.Group(children),
	)
}

func ImagePhotoCompsitionHorizontalCameraHorizontalPanoramaCompositionPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.81 11.2a.52.52 0 0 0 .47 0a.49.49 0 0 0 .22-.41V3.26a.49.49 0 0 0-.22-.41a.52.52 0 0 0-.47-.05a14.67 14.67 0 0 1-11.62 0a.52.52 0 0 0-.47.05a.49.49 0 0 0-.22.41v7.48a.49.49 0 0 0 .22.41a.52.52 0 0 0 .47 0a14.67 14.67 0 0 1 11.62.05Z"/>`), g.Group(children),
	)
}

func ImagePhotoFocusFramePhotosFramePhotoCameraPhotographyPicturesFocus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx="1"/><path d="M11 6.5V9H8.5M3 7.5V5h2.5"/></g>`), g.Group(children),
	)
}

func ImagePhotoFourPhotosCameraPicturePhotographyPicturesFourPhoto(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="5" x=".5" y=".5" rx=".5"/><rect width="5" height="5" x="8.5" y=".5" rx=".5"/><rect width="5" height="5" x=".5" y="8.5" rx=".5"/><rect width="5" height="5" x="8.5" y="8.5" rx=".5"/></g>`), g.Group(children),
	)
}

func ImagePhotoPolaroidFourPhotosCameraPolaroidPicturePhotographyPicturesFourPhoto(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5.25" height="5.25" x=".5" y=".5" rx=".5"/><rect width="5.25" height="5.25" x="8.25" y=".5" rx=".5"/><rect width="5.25" height="5.25" x=".5" y="8.25" rx=".5"/><rect width="5.25" height="5.25" x="8.25" y="8.25" rx=".5"/><path d="M.5 3.5h5.25m2.5 0h5.25m-5.25 7.75h5.25m-13 0h5.25"/></g>`), g.Group(children),
	)
}

func ImagePhotoPolaroidPhotosPolaroidPictureCameraPhotographyPhotoPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="9" x=".5" y=".5" rx="1" transform="rotate(180 5 5)"/><path d="m12 4.7l.78.25a1 1 0 0 1 .64 1.27l-2.18 6.59a1 1 0 0 1-1.26.64L5.56 12M.5 7h9"/></g>`), g.Group(children),
	)
}

func ImagePictureFlowerPhotosPhotoPictureCameraPhotographyPicturesFlower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M10 6.5a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3V4.31a.5.5 0 0 1 .72-.45L7 5l2.28-1.14a.5.5 0 0 1 .72.45Zm-3 3v4m0 0L9.5 11M7 13.5L4.5 11"/></g>`), g.Group(children),
	)
}

func ImagePictureGalleryPagesFilterPicturePaginationImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="8.5" x="3" y="4" rx="1" transform="rotate(180 8.25 8.25)"/><path d="M.5 10V2.5a1 1 0 0 1 1-1h9M3.6 12.42l3.93-4.15A1 1 0 0 1 9 8.26l3.95 4.14"/></g>`), g.Group(children),
	)
}

func ImagePictureLandscapeOnePhotosPhotoLandscapePicturePhotographyCameraPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7.005 7)"/><circle cx="9.25" cy="4.75" r="1.75"/><path d="M9.4 13.5a7.36 7.36 0 0 0-7.4-6a7.88 7.88 0 0 0-1.5.14"/><path d="M13.5 9.91A7.8 7.8 0 0 0 11 9.5a7.89 7.89 0 0 0-3.13.64"/></g>`), g.Group(children),
	)
}

func ImagePictureLandscapeTwoPhotosPhotoLandscapePicturePhotographyCameraPictures(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M.5 11h13m-9.66 0l5.21-4.88a.5.5 0 0 1 .64 0l3.81 2.73"/><circle cx="4.5" cy="4.5" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceAddCircleButtonRemoveCrossAddButtonsPlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 4v6M4 7h6"/></g>`), g.Group(children),
	)
}

func InterfaceAddOneExpandCrossButtonsButtonMoreRemovePlusAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .54v13M.5 7h13"/>`), g.Group(children),
	)
}

func InterfaceAddSquareSquareRemoveCrossButtonsAddPlusButton(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4v6M4 7h6"/><rect width="13" height="13" x=".5" y=".5" rx="3"/></g>`), g.Group(children),
	)
}

func InterfaceAddTwoRemoveBoldCrossButtonsButtonAddPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.5a1 1 0 0 0-1-1h-4v-4a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v4h-4a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h4v4a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-4h4a1 1 0 0 0 1-1Z"/>`), g.Group(children),
	)
}

func InterfaceAlertAlarmBellOffDisableSilentNotificationOffSilenceAlarmBellAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13M6 13.5h2M8.73.84A4.51 4.51 0 0 0 2.5 5v2.5M3 11h10.5a2 2 0 0 1-2-2V5a4.42 4.42 0 0 0-.5-2"/>`), g.Group(children),
	)
}

func InterfaceAlertAlarmBellOneNotificationVibrateRingSoundAlarmAlertBellNoise(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 13.25h2m3-7.5a4 4 0 0 0-8 0v3.5a1.5 1.5 0 0 1-1.5 1.5h11a1.5 1.5 0 0 1-1.5-1.5ZM.5 5.62A6 6 0 0 1 3 .75m10.5 4.87A6 6 0 0 0 11 .75"/>`), g.Group(children),
	)
}

func InterfaceAlertAlarmBellTwoAlertBellRingNotificationAlarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5a4.29 4.29 0 0 1 4.29 4.29c0 4.77 1.74 5.71 2.21 5.71H.5c.48 0 2.21-.95 2.21-5.71A4.29 4.29 0 0 1 7 .5ZM5.5 12.33a1.55 1.55 0 0 0 3 0"/>`), g.Group(children),
	)
}

func InterfaceAlertInformationCircleInformationFrameInfoMoreHelpPointCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 7v3.5"/><circle cx="7" cy="4.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceAlertRadioActiveOneDangerNukeRadiationNuclearWarningAlertRadioactiveCaution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 4.83L3.75.94A6.49 6.49 0 0 0 .5 6.56H5a2 2 0 0 1 1-1.73Zm3 1.73h4.5A6.49 6.49 0 0 0 10.25.94L8 4.83a2 2 0 0 1 1 1.73Zm-2 2a2 2 0 0 1-1-.26l-2.25 3.89a6.51 6.51 0 0 0 6.5 0L8 8.3a2 2 0 0 1-1 .26Z"/>`), g.Group(children),
	)
}

func InterfaceAlertRadioActiveThreeWarningRadioactiveRadiationEmergencyDangerSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7.49" r="3.5"/><path d="M9.35 1.31a3.32 3.32 0 1 1-4.7 0"/><path d="M.5 8.5a3.32 3.32 0 1 1 2.35 4.07"/><path d="M11.15 12.57A3.32 3.32 0 1 1 13.5 8.5"/></g>`), g.Group(children),
	)
}

func InterfaceAlertRadioActiveTwoWarningRadioactiveRadiationEmergencyDangerSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.92 1.25a3.2 3.2 0 1 1-3.84 0"/><path d="M.5 9.17a3.2 3.2 0 1 1 1.92 3.32"/><path d="M11.58 12.49a3.19 3.19 0 1 1 1.92-3.32"/></g>`), g.Group(children),
	)
}

func InterfaceAlertWarningCircleWarningAlertFrameExclamationCautionCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 3.5v3"/><circle cx="7" cy="9.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceAlertWarningDiamondDiamondAlertWarningFrameExclamationCaution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 3.5v3"/><circle cx="7" cy="9.5" r=".5"/><rect width="9.82" height="9.82" x="2.09" y="2.09" rx="1.07" transform="rotate(-45 7 7)"/></g>`), g.Group(children),
	)
}

func InterfaceAlertWarningTriangleFrameAlertWarningTriangleExclamationCaution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 5v3"/><circle cx="7" cy="11" r=".5"/><path d="M7.89 1.05a1 1 0 0 0-1.78 0l-5.5 11a1 1 0 0 0 .89 1.45h11a1 1 0 0 0 .89-1.45Z"/></g>`), g.Group(children),
	)
}

func InterfaceAlignBackBackDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="10.5" x=".5" y=".5" rx="1"/><path d="M13.5 3.5v9a1 1 0 0 1-1 1h-9"/></g>`), g.Group(children),
	)
}

func InterfaceAlignFrontDesignFrontLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="10.5" x="3" y="3" rx="1" transform="rotate(180 8.25 8.25)"/><path d="M.5 10.5v-9a1 1 0 0 1 1-1h9"/></g>`), g.Group(children),
	)
}

func InterfaceAlignHorizontalCenterAlignCenterDesign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="4" x="3.5" y=".5" rx="1"/><rect width="11" height="4" x="1.5" y="9.5" rx="1"/><path d="M7 4.5v5"/></g>`), g.Group(children),
	)
}

func InterfaceAlignHorizontalLeftAlignDesignLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13"/><rect width="6.5" height="4" x="3" y="1.5" rx="1"/><rect width="10.5" height="4" x="3" y="8.5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceAlignHorizontalRightAlignDesignRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5v13"/><rect width="6.5" height="4" x="4.5" y="1.5" rx="1" transform="rotate(180 7.75 3.5)"/><rect width="10.5" height="4" x=".5" y="8.5" rx="1" transform="rotate(180 5.75 10.5)"/></g>`), g.Group(children),
	)
}

func InterfaceAlignLayersOneDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.47 6.9a1.18 1.18 0 0 1-.94 0L.83 4.26a.56.56 0 0 1 0-1L6.53.6a1.18 1.18 0 0 1 .94 0l5.7 2.64a.56.56 0 0 1 0 1Zm6.03.45l-6.1 2.81a1 1 0 0 1-.83 0L.5 7.35m13 3.25l-6.1 2.81a1 1 0 0 1-.83 0L.5 10.6"/>`), g.Group(children),
	)
}

func InterfaceAlignLayersTwoDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="6" x="5.5" y="1.5" rx="1"/><path d="M11 10H4a1 1 0 0 1-1-1V4"/><path d="M9 12.5H1.5a1 1 0 0 1-1-1V6"/></g>`), g.Group(children),
	)
}

func InterfaceAlignVerticalBottomAlignDesignBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 13.5H.5"/><rect width="6.5" height="4" x="7.25" y="5.75" rx="1" transform="rotate(-90 10.5 7.75)"/><rect width="10.5" height="4" x="-1.75" y="3.75" rx="1" transform="rotate(-90 3.5 5.75)"/></g>`), g.Group(children),
	)
}

func InterfaceAlignVerticalCenterAlignDesignMiddle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="4" x="7.5" y="5.5" rx="1" transform="rotate(90 11.5 7.5)"/><rect width="11" height="4" x="-3" y="5" rx="1" transform="rotate(90 2.5 7)"/><path d="M9.5 7h-5"/></g>`), g.Group(children),
	)
}

func InterfaceAlignVerticalTopAlignDesignTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5H.5"/><rect width="6.5" height="4" x="7.25" y="4.25" rx="1" transform="rotate(90 10.5 6.25)"/><rect width="10.5" height="4" x="-1.75" y="6.25" rx="1" transform="rotate(90 3.5 8.25)"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsAllDirectionMoveButtonArrowsDirection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 2.5l2-2l2 2m-4 9l2 2l2-2M7 .5v13M11.5 5l2 2l-2 2m-9-4l-2 2l2 2m11-2H.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsBendDownLeftOneArrowBendCurveChangeDirectionLeftToDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.75 11l2.5 2.5l2.5-2.5"/><path d="M11.25.5h-5a1 1 0 0 0-1 1v12"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendDownLeftThreeArrowBendCurveChangeDirectionLeftToDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 10.5l3 3l3-3"/><path d="M11.5.5h-2a4 4 0 0 0-4 4v9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendDownLeftTwoArrowBendCurveChangeDirectionDownToLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 11.5l-3-3l3-3"/><path d="M13.5 2.5v2a4 4 0 0 1-4 4h-9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendDownRightOneArrowBendCurveChangeDirectionRightToDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 10.5l-3 3l-3-3"/><path d="M2.5.5h2a4 4 0 0 1 4 4v9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendDownRightTwoArrowBendCurveChangeDirectionDownToRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 11.5l3-3l-3-3"/><path d="M.5 2.5v2a4 4 0 0 0 4 4h9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendLeftOneArrowBendCurveChangeDirectionUpToLeftBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 2.75L.5 5.25L3 7.75"/><path d="M13.5 11.25v-5a1 1 0 0 0-1-1H.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendRightOneArrowBendCurveChangeDirectionUpToRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11 2.75l2.5 2.5l-2.5 2.5"/><path d="M.5 11.25v-5a1 1 0 0 1 1-1h12"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendUpLeftOneArrowBendCurveChangeDirectionLeftToUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.75 3L5.25.5L7.75 3"/><path d="M11.25 13.5h-5a1 1 0 0 1-1-1V.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendUpLeftThreeArrowBendCurveChangeDirectionUpToLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 2.5l-3 3l3 3"/><path d="M13.5 11.5v-2a4 4 0 0 0-4-4h-9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendUpLeftTwoArrowBendCurveChangeDirectionLeftToUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 3.5l3-3l3 3"/><path d="M11.5 13.5h-2a4 4 0 0 1-4-4v-9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendUpRightOneArrowBendCurveChangeDirectionRightToUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 3.5l-3-3l-3 3"/><path d="M2.5 13.5h2a4 4 0 0 0 4-4v-9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBendUpRightTwoArrowBendCurveChangeDirectionUpToRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 2.5l3 3l-3 3"/><path d="M.5 11.5v-2a4 4 0 0 1 4-4h9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsBottomDownMoveArrowArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4 6.5l3 3l3-3m-3-6v9m-3.5 4h7"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonDownArrowDownKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 3.85L6.65 10a.48.48 0 0 0 .7 0l6.15-6.15"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonDownDoubleArrowArrowsDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 6.46l6.15 6.14a.48.48 0 0 0 .7 0l6.15-6.14"/><path d="M.5 1.25L6.65 7.4a.5.5 0 0 0 .7 0l6.15-6.15"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsButtonLeftArrowKeyboardLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.15.5L4 6.65a.48.48 0 0 0 0 .7l6.15 6.15"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonLeftDoubleArrowArrowsDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.54.5L1.4 6.65a.48.48 0 0 0 0 .7l6.14 6.15"/><path d="M12.75.5L6.6 6.65a.5.5 0 0 0 0 .7l6.15 6.15"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsButtonRightArrowRightKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.85.5L10 6.65a.48.48 0 0 1 0 .7L3.85 13.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonRightDuobleArrowArrowsDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6.46.5l6.14 6.15a.48.48 0 0 1 0 .7L6.46 13.5"/><path d="M1.25.5L7.4 6.65a.5.5 0 0 1 0 .7L1.25 13.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsButtonToBottomArrowDownLineToBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 5.71l6.15 6.14a.48.48 0 0 0 .7 0l6.15-6.14M.5 2h13"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonToLeftArrowLineToLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.29.5L2.15 6.65a.48.48 0 0 0 0 .7l6.14 6.15M12 .5v13"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonToRightArrowLineToRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.71.5l6.14 6.15a.48.48 0 0 1 0 .7L5.71 13.5M2 .5v13"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonToTopArrowUpLineToTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 8.29l6.15-6.14a.48.48 0 0 1 .7 0l6.15 6.14M.5 12h13"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonUpArrowUpKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 10.15L6.65 4a.48.48 0 0 1 .7 0l6.15 6.15"/>`), g.Group(children),
	)
}

func InterfaceArrowsButtonUpDoubleArrowArrowsDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 7.54L6.65 1.4a.48.48 0 0 1 .7 0l6.15 6.14"/><path d="M.5 12.75L6.65 6.6a.5.5 0 0 1 .7 0l6.15 6.15"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsButtonZigzagBothDirectionArrowCurvyDiagramZigzagVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 11L3 13.5v-11a2 2 0 0 1 4 0v9a2 2 0 0 0 4 0V.5L13.5 3"/>`), g.Group(children),
	)
}

func InterfaceArrowsCornerDownLeftDownKeyboardArrowLeftBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5.5l-13 13m5 0h-5v-5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCornerDownRightDownKeyboardArrowRightBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13m-5 0h5v-5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCornerUpLeftKeyboardTopArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5L.5.5m5 0h-5v5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCornerUpRightKeyboardTopArrowRightUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13m-5 0h5v5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCrossOverCrossMoveOverArrowArrowsIght(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13m-4 0h4v4M9 9l4.5 4.5m-4 0h4v-4M5 5L.5.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCrossoverDownCrossMoveOverArrowArrowsDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13m0-4v4h-4M5 9L.5 13.5m0-4v4h4M9 5L13.5.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCrossoverLeftCrossMoveOverArrowArrowsLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5.5l-13 13m4 0h-4v-4M5 5L.5.5m4 0h-4v4M9 9l4.5 4.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCrossoverUpCrossMoveOverArrowArrowsRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5L.5.5m0 4v-4h4M9 5L13.5.5m0 4v-4h-4M5 9L.5 13.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCurvyBothDirectionOneBothDirectionArrowCurvyDiagramZigzagHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 .5L13.5 3h-12a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h11a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H.5L3 13.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCurvyBothDirectionTwoBothDirectionArrowCurvyDiagramZigzagHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 .5L.5 3h12a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-11a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h12L11 13.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsCurvyLeftSnakeArrowSidewaysDiagramLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.75h-12a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H.5"/><path d="m3 8.25l-2.5 2.5l2.5 2.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsCurvyRightSnakeArrowSidewaysDiagramRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.75h12a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-11a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h12"/><path d="m11 8.25l2.5 2.5l-2.5 2.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDataTransferDiagonalArrowsArrowServerDataDiagonalInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 10.5l10-10H6m7.5 3l-10 10H8"/>`), g.Group(children),
	)
}

func InterfaceArrowsDataTransferDiagonalSquareArrowSquareDataDiagonalInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m4 8l4.5-4.5H5M10 6l-4.5 4.5H9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDataTransferHorizontalArrowSquareDataInternetTransferNetworkHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m6 10.5l-2-2h6m-2-5l2 2H4"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDataTransferVerticalArrowSquareDataInternetTransferNetworkVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m10.5 8l-2 2V4m-5 2l2-2v6"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDataTrasnferVerticalArrowServerArrowsDataVerticalInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 3.5l3-3v13m6-3l-3 3V.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsDiagonalArrowsArrowServerDataDiagonalInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.5 8.5l8-8m0 3.5V.5h0H6m6.5 5l-8 8m0-3.5v3.5H8"/>`), g.Group(children),
	)
}

func InterfaceArrowsDiagonalOneExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13m-4 0h4v4m-9 9h-4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsDiagonalScrollPointOneMoveScroll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="1.5"/><path d="M9 .5h4a.5.5 0 0 1 .5.5v4M5 13.5H1a.5.5 0 0 1-.5-.5V9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDiagonalScrollPointTwoMoveScroll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="1.5"/><path d="M5 .5H1a.5.5 0 0 0-.5.5v4M9 13.5h4a.5.5 0 0 0 .5-.5V9"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDiagonalTwoExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5.5h-4v4m0-4l13 13m-4 0h4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsDownArrowDownKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v13m3.5-3.5L7 13.5L3.5 10"/>`), g.Group(children),
	)
}

func InterfaceArrowsDownCircleOneArrowKeyboardCircleButtonDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4v6M5.5 8.5L7 10l1.5-1.5"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDownCircleTwoArrowKeyboardCircleButtonDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 6L7 9L4 6"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsDownleftCornerArrowCornerDownLeftDownleft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5h-12a1 1 0 0 1-1-1V.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsDownrightCornerArrowCornerDownRightDownright(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5v12a1 1 0 0 1-1 1H.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandDiagonalOneExpandResizeBiggerDiagonalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 5L13.5.5m-4 0h4v4M5 9L.5 13.5m4 0h-4v-4M4 4l6 6"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandDiagonalTwoExpandResizeBiggerDiagonalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 5L.5.5m4 0h-4v4M9 9l4.5 4.5m-4 0h4v-4M10 4l-6 6"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandFiveExpandSmallBiggerRetractSmallerBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 8v4.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1H6m4 0h3.5V4m0-3.5L7 7"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandFourExpandSmallBiggerRetractSmallerBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x="3.5" y=".5" rx="1" transform="rotate(180 8.5 5.5)"/><path d="M10.5 13.5h-9a1 1 0 0 1-1-1v-9m7 0h3v3m0-3l-4 4"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsExpandOneExpandSmallBiggerRetractSmallerBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 5.5l5-5m-4 0h4v4m-8 1l-5-5m4 0h-4v4m8 4l5 5m-4 0h4v-4m-8-1l-5 5m4 0h-4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandThreeExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 5.5l5-5m-4 0h4v4m-8 4l-5 5m4 0h-4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandTwoExpandSmallerRetractBiggerBigSmallDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 5.5l-5-5m4 0h-4v4m8 4l5 5m-4 0h4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsExpandVerticalDiamondArrowDiamondDataVerticalDataInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9.82" height="9.82" x="2.09" y="2.09" rx="1.07" transform="rotate(-45 7 7)"/><path d="M5.5 5.5L7 4l1.5 1.5m-3 3L7 10l1.5-1.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsHorizonalScrollPointMoveScrollHorizonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.48 3.87L.61 6.74a.36.36 0 0 0 0 .52l2.87 2.87m7.04-6.26l2.87 2.87a.36.36 0 0 1 0 .52l-2.87 2.87"/><circle cx="7" cy="7" r="1.25"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsHorizontalExpandExpandResizeBiggerHorizontalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 7h-9M5 4.5L2.5 7L5 9.5m4-5L11.5 7L9 9.5M.5.5v13m13-13v13"/>`), g.Group(children),
	)
}

func InterfaceArrowsHorizontalExpandResizeBiggerHorizontalSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7H.5m3-3l-3 3l3 3m7-6l3 3l-3 3"/>`), g.Group(children),
	)
}

func InterfaceArrowsHorizontalShrinkResizeShrinkBiggerHorizontalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 4.5L8.5 7L11 9.5m-8-5L5.5 7L3 9.5M.5.5v13m13-13v13"/>`), g.Group(children),
	)
}

func InterfaceArrowsHorizontalUpDownExpandResizeBiggerVerticalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.13 3.48L7.26.61a.36.36 0 0 0-.52 0L3.87 3.48m6.26 7.04l-2.87 2.87a.36.36 0 0 1-.52 0l-2.87-2.87M.5 7h13"/>`), g.Group(children),
	)
}

func InterfaceArrowsLeftArrowKeyboardLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7H.5M4 3.5L.5 7L4 10.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsLeftCircleOneArrowKeyboardCircleButtonLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 7H4m1.5-1.5L4 7l1.5 1.5"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsLeftCircleTwoArrowKeyboardCircleButtonLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 4L5 7l3 3"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsLoopArrowDiagramLoopInfinityRepeat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 10.43a2.37 2.37 0 0 1-1.5.57a2.75 2.75 0 0 1-3-3a2.75 2.75 0 0 1 3-3c2.75 0 4.25 6 7 6a2.75 2.75 0 0 0 3-3a2.75 2.75 0 0 0-3-3h-2l2-2"/>`), g.Group(children),
	)
}

func InterfaceArrowsMoveDownMoveDownArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 7.5v6M4.5 11L7 13.5L9.5 11"/><rect width="4.5" height="13" x="4.75" y="-3.75" rx="1" transform="rotate(-90 7 2.75)"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveHorizontalCircleTransferDataInternetArrowHorizontalNetworkCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 10.5l-2-2h6m-2-5l2 2H4"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveLeftDownExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5.5l-2 2l2 2"/><path d="M11.5 13.5v-10a1 1 0 0 0-1-1H.5"/><path d="m9.5 11.5l2 2l2-2"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveLeftMoveLeftArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7h-6M3 4.5L.5 7L3 9.5"/><rect width="4.5" height="13" x="9" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveLeftUpExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 13.5l-2-2l2-2"/><path d="M11.5.5v10a1 1 0 0 1-1 1H.5"/><path d="m9.5 2.5l2-2l2 2"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveRightDownExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5.5l2 2l-2 2"/><path d="M2.5 13.5v-10a1 1 0 0 1 1-1h10"/><path d="m4.5 11.5l-2 2l-2-2"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveRightMoveRightArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 7h6M11 4.5L13.5 7L11 9.5"/><rect width="4.5" height="13" x=".5" y=".5" rx="1" transform="rotate(-180 2.75 7)"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveRightUpExpandResizeBiggerCornerSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 2.5l2-2l2 2"/><path d="M13.5 11.5h-10a1 1 0 0 1-1-1V.5"/><path d="m11.5 9.5l2 2l-2 2"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsMoveUpMoveUpArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 6.5v-6M9.5 3L7 .5L4.5 3"/><rect width="4.5" height="13" x="4.75" y="4.75" rx="1" transform="rotate(90 7 11.25)"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsReloadOneArrowsLoadArrowSyncSquareLoadingReloadSynchronize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 5L3 2.5L5.5 5"/><path d="M6 13.5H4a1 1 0 0 1-1-1v-10M13.5 9L11 11.5L8.5 9"/><path d="M8 .5h2a1 1 0 0 1 1 1v10"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsReloadTwoArrowsLoadArrowSyncSquareLoadingReloadSynchronize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9 .5L11.5 3L9 5.5"/><path d="M.5 6V4a1 1 0 0 1 1-1h10M5 13.5L2.5 11L5 8.5"/><path d="M13.5 8v2a1 1 0 0 1-1 1h-10"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsRightArrowRightKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7h13M10 10.5L13.5 7L10 3.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsRightCircleOneArrowKeyboardCircleButtonRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 7h6M8.5 5.5L10 7L8.5 8.5"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsRightCircleTwoArrowKeyboardCircleButtonRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 4l3 3l-3 3"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsRoundLeftDiagramRoundArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.75.5L2.25 3l2.5 2.5"/><path d="M1.75 8.25A5.25 5.25 0 1 0 7 3H2.25"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsRoundRightDiagramRoundArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9.25.5l2.5 2.5l-2.5 2.5"/><path d="M12.25 8.25A5.25 5.25 0 1 1 7 3h4.75"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkFiveExpandRetractShrinkBiggerBigSmallSmaller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 6V1.5a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H8m-4 0H.5V10m0 3.5L7 7"/>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkFourExpandRetractShrinkBiggerBigSmallSmaller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x="3.5" y=".5" rx="1" transform="rotate(180 8.5 5.5)"/><path d="M10.5 13.5h-9a1 1 0 0 1-1-1v-9m9 4h-3v-3m0 3l4-4"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkOneExpandRetractShrinkBiggerBigSmallSmaller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l4-4M1 9.5h3.5V13m9 .5l-4-4m3.5 0H9.5V13M.5.5l4 4M1 4.5h3.5V1m9-.5l-4 4m3.5 0H9.5V1"/>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkThreeExpandRetractShrinkBiggerBigSmallSmaller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l5-5m-4 0h4v4m8-12l-5 5m4 0h-4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkTwoExpandRetractShrinkBiggerBigSmallSmaller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 13.5l-5-5m4 0h-4v4M.5.5l5 5m-4 0h4v-4"/>`), g.Group(children),
	)
}

func InterfaceArrowsShrinkVerticalMoveVerticalShrink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.25 13.5L6.74 9a.37.37 0 0 1 .52 0l4.49 4.49M2.25.5L6.74 5a.37.37 0 0 0 .52 0L11.75.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsShuffleMultimediaShuffleMultiButtonControlsMedia(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 2l2 2l-2 2m-3-2h5m-2 4l2 2l-2 2"/><path d="M.5 4H4l4 6h5.5m-13 0H4"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsSplitVerticalUpMergeArrowDiagram(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 3.5l-3-3l-3 3m7 10a4 4 0 0 1-4-4a4 4 0 0 1-4 4m4-4v-9"/>`), g.Group(children),
	)
}

func InterfaceArrowsSynchronizeArrowsLoadingLoadSyncSynchronizeArrowReload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11 9l2-.5l.5 2"/><path d="M13 8.5A6.76 6.76 0 0 1 7 13h0a6 6 0 0 1-5.64-3.95M3 5l-2 .5l-.5-2"/><path d="M1 5.5C1.84 3.2 4.42 1 7 1h0a6 6 0 0 1 5.64 4"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsSynchronizeWarningArrowFailNotificationSyncWarningFailureSynchronizeError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11 9l2-.5l.5 2"/><path d="M13 8.5A6.6 6.6 0 0 1 7 13h0a6 6 0 0 1-5.64-3.95M3 5l-2 .5l-.5-2"/><path d="M1 5.5A6.79 6.79 0 0 1 7 1h0a6 6 0 0 1 5.64 4M7 3.5v4"/><circle cx="7" cy="10" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsTopMoveUpArrowArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4 7.5l3-3l3 3m-3 6v-9M3.5.5h7"/>`), g.Group(children),
	)
}

func InterfaceArrowsTriangleLoopDiagramTriangleLoopArrowArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 13.5L4 11h8.5m1-7l-1 3.5L8 .5m-6 3h3.5l-5 7"/>`), g.Group(children),
	)
}

func InterfaceArrowsTurnBackwardArrowBendCurveChangeDirectionReturnLeftBackBackward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 1.5l-3 3l3 3"/><path d="M.5 4.5h9a4 4 0 0 1 0 8h-5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsTurnDownArrowBendCurveChangeDirectionReturnDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.5 10.5l3 3l3-3"/><path d="M4.5 13.5v-9a4 4 0 0 1 8 0v5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsTurnForwardArrowBendCurveChangeDirectionReturnRightNextForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 1.5l3 3l-3 3"/><path d="M13.5 4.5h-9a4 4 0 0 0 0 8h5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsTurnUpArrowBendCurveChangeDirectionReturnUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12.5 3.5l-3-3l-3 3"/><path d="M9.5.5v9a4 4 0 0 1-8 0v-5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsUpArrowUpKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V.5M10.5 4L7 .5L3.5 4"/>`), g.Group(children),
	)
}

func InterfaceArrowsUpCircleOneArrowUpKeyboardCircleButton(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 10V4M5.5 5.5L7 4l1.5 1.5"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsUpCircleTwoArrowKeyboardCircleButtonUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4 8l3-3l3 3"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceArrowsUpleftCornerArrowCornerUpLeftUpleft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5h-12a1 1 0 0 0-1 1v12"/>`), g.Group(children),
	)
}

func InterfaceArrowsUprightCornerArrowUpRightUprightCorner(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5h12a1 1 0 0 1 1 1v12"/>`), g.Group(children),
	)
}

func InterfaceArrowsVerticalExpandOneMoveExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.25 5.1L6.74.61a.36.36 0 0 1 .52 0l4.49 4.49m-9.5 3.8l4.49 4.49a.36.36 0 0 0 .52 0l4.49-4.49"/>`), g.Group(children),
	)
}

func InterfaceArrowsVerticalExpandResizeBiggerVerticalSmallSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V.5m3 3l-3-3l-3 3m6 7l-3 3l-3-3"/>`), g.Group(children),
	)
}

func InterfaceArrowsVerticalExpandTwoExpandResizeBiggerVerticalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 11.5v-9M9.5 5L7 2.5L4.5 5m5 4L7 11.5L4.5 9m9-8.5H.5m13 13H.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsVerticalLeftRightExpandResizeBiggerHorizontalSmallerSizeArrowArrowsBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.48 3.87L.61 6.74a.36.36 0 0 0 0 .52l2.87 2.87m7.04-6.26l2.87 2.87a.36.36 0 0 1 0 .52l-2.87 2.87M7 13.5V.5"/>`), g.Group(children),
	)
}

func InterfaceArrowsVerticalScrollPointMoveScrollVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.13 3.48L7.26.61a.36.36 0 0 0-.52 0L3.87 3.48m6.26 7.04l-2.87 2.87a.36.36 0 0 1-.52 0l-2.87-2.87"/><circle cx="7" cy="7" r="1.25"/></g>`), g.Group(children),
	)
}

func InterfaceAwardCrownRewardSocialRatingMediaQueenVipKingCrown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 4l-3 3L7 2L3.5 7l-3-3v6.5A1.5 1.5 0 0 0 2 12h10a1.5 1.5 0 0 0 1.5-1.5Z"/>`), g.Group(children),
	)
}

func InterfaceAwardHalfStarRewardRatingRateSocialStarMediaFavoriteLikeStarsHalf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25.5a.54.54 0 0 0-.49.32L8.17 4.18a.52.52 0 0 1-.41.31L4.22 5a.58.58 0 0 0-.3 1l2.56 2.63a.58.58 0 0 1 .16.5L6 12.83a.56.56 0 0 0 .8.6l3.2-1.74a.59.59 0 0 1 .26-.07Z"/>`), g.Group(children),
	)
}

func InterfaceAwardTrophyRewardRatingTrophySocialAwardMedia(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 8.5v5m-2.5 0h5m-6-8a3 3 0 0 1-3-3v-1H4v4Zm7 0a3 3 0 0 0 3-3v-1H10v4Z"/><path d="M10 5.5a3 3 0 0 1-6 0v-5h6Z"/></g>`), g.Group(children),
	)
}

func InterfaceBlockRemoveCircleGarbageTrashDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m2.4 11.6l9.2-9.2"/></g>`), g.Group(children),
	)
}

func InterfaceBookmarkBookmarksTagsFavorite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11 13.5l-4-4l-4 4v-12a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Z"/>`), g.Group(children),
	)
}

func InterfaceBookmarkDoubleBookmarksDoubleTagsFavorite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9 13.5l-3.5-3l-3.5 3v-9a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1Z"/><path d="M5 .5h6a1 1 0 0 1 1 1v9"/></g>`), g.Group(children),
	)
}

func InterfaceBookmarkTagTagsBookmarkFavorite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.28 9.39l-3.89 3.89a.75.75 0 0 1-1.06 0L.61 5.56a.36.36 0 0 1-.11-.29l.59-3.83a.37.37 0 0 1 .35-.35L5.27.5a.36.36 0 0 1 .29.11l7.72 7.72a.75.75 0 0 1 0 1.06Z"/><circle cx="4.11" cy="4.11" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarAddAddCalendarDateDayMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5m1 5.5h-5M7 5.5v5"/>`), g.Group(children),
	)
}

func InterfaceCalendarBlankCalendarDateDayMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-10 4h13m-10-6v4m7-4v4m-7-2h5"/>`), g.Group(children),
	)
}

func InterfaceCalendarBlockBlockCalendarDateDayMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><circle cx="7" cy="8.5" r="3"/><path d="m4.88 10.62l4.24-4.24"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarCheckApproveCalendarCheckDateDayMonthSuccess(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="m4 9l2 1.5l3.5-4"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarDateMonthThirtyThirtyCalendarDateWeekDayMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 6.5h1V11m-1 0h2"/><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M3.5 6.5H6l-1.5 2s1.5 0 1.5 1A1.33 1.33 0 0 1 4.5 11h-1"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarDateOneOneCalendarDateDayMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6 6.5h1V11m-1.5 0h3"/><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarDisableCalendarDateDayDeleteDisableMonthRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5m.27 7.27L5.23 6.23m3.54 0L5.23 9.77"/>`), g.Group(children),
	)
}

func InterfaceCalendarDownloadArrowCalendarDateDayDownDownloadMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="m5 9l2 2l2-2m-2 2V6"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarFavoriteCalendarDateDayFavoriteLikeMonthStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M6.45 5.37a.59.59 0 0 1 1.1 0L8.19 7H9.9a.6.6 0 0 1 .4 1L8.79 9.36l.64 1.28a.58.58 0 0 1-.12.69a.61.61 0 0 1-.7.1L7 10.55l-1.61.88a.61.61 0 0 1-.7-.1a.58.58 0 0 1-.12-.69l.64-1.28L3.7 8a.6.6 0 0 1 .4-1h1.71Z"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarHeartCalendarDateDayFavoriteHeartLikeMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M7 7.5c1-2 3-1 3 .5c0 2-3 3-3 3s-3-1-3-3c0-1.5 2-2.5 3-.5Z"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarHomeCalendarDateDayHomeHouseMonth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="m10 8.5l-3-3l-3 3"/><path d="M5 7.5V11h4V7.5"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarLockCalendarDateDayLockMonthPadlockSecureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M5 8h4v3H5zm.5 0V7a1.5 1.5 0 0 1 3 0v1"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarMarkCalendarDateDayMonthMark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><circle cx="3.5" cy="7.5" r=".5"/><circle cx="7" cy="7.5" r=".5"/><circle cx="10.5" cy="7.5" r=".5"/><circle cx="3.5" cy="10.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarRemoveCalendarDateDayDeleteMonthRemoveSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5m1 7.5h-5m-4-3.5h13"/>`), g.Group(children),
	)
}

func InterfaceCalendarSettingCalendarCogDateDayGearLoadLoadingMonthSettingWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5m-1.53 2V6m-3.03.25l1.3.75m-1.3 2.75L5.24 9m1.73 2.5V10M10 9.75L8.7 9M10 6.25L8.7 7"/><circle cx="6.97" cy="8" r="2"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarUploadCalendarDateDayMonthPushUpArrowUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="m5 8l2-2l2 2M7 6v5"/></g>`), g.Group(children),
	)
}

func InterfaceCalendarWarningAlterCalendarCautionDateDayMonthNotificationWarning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5M7 5v3"/><circle cx="7" cy="11" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceContentArchiveBoxContentBankerArchiveFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 5.54h11v7a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-7h0Z"/><rect width="4" height="13" x="5" y="-2.96" rx="1" transform="rotate(90 7 3.54)"/><path d="M5.5 8.54h3"/></g>`), g.Group(children),
	)
}

func InterfaceContentArchiveFolderOutboxContentFolderArchiveFileInbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 9.54a1 1 0 0 0-1-1h-3v1h-5v-1h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1Zm-12-4h11m-11-3h11"/>`), g.Group(children),
	)
}

func InterfaceContentArchiveLockerLockerContentArchiveFileCabinet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5.5h11a1 1 0 0 1 1 1v10a2 2 0 0 1-2 2h-9a2 2 0 0 1-2-2v-10a1 1 0 0 1 1-1ZM.5 7h13"/><circle cx="7" cy="3.5" r=".5"/><circle cx="7" cy="10" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceContentBookContentBooksBookClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 13.54H3a1.5 1.5 0 0 1 0-3h8.5a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1H3A1.5 1.5 0 0 0 1.5 2v10m10-1.46v3"/>`), g.Group(children),
	)
}

func InterfaceContentBookEditPencilContentWriteNotebookBookEditCompositionCreation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 13.5H1.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1H9m1.5 3l1.5-3l1.5 3V12a1.5 1.5 0 0 1-3 0Zm0 6h3m-10-9v13M6 4h2"/>`), g.Group(children),
	)
}

func InterfaceContentBookOpenContentBooksBookOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5a9.26 9.26 0 0 0-5.61-2.95a1 1 0 0 1-.89-1V1.5A1 1 0 0 1 .85.74a1 1 0 0 1 .79-.23A9.3 9.3 0 0 1 7 3.43Zm0 0a9.26 9.26 0 0 1 5.61-2.95a1 1 0 0 0 .89-1V1.5a1 1 0 0 0-.35-.76a1 1 0 0 0-.79-.23A9.3 9.3 0 0 0 7 3.43Z"/>`), g.Group(children),
	)
}

func InterfaceContentBookPagePagesContentBooksBookOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 12.5a1 1 0 0 1-1.45.89l-4-2A1 1 0 0 1 2 10.5V.5l5.45 2.72a1 1 0 0 1 .55.9Z"/><path d="M2 .5h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H8"/></g>`), g.Group(children),
	)
}

func InterfaceContentBookTwoLibraryContentBooksBookShelfStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3.5" height="13" x=".55" y=".5" rx=".5"/><rect width="3.5" height="11" x="4.05" y="2.5" rx=".5"/><rect width="3" height="11" x="9.26" y="2.3" rx=".5" transform="rotate(-14.05 10.779 7.795)"/><path d="M.55 10h3.5m0-1h3.5m2.5 2l2.88-.72"/></g>`), g.Group(children),
	)
}

func InterfaceContentChartProductDataAnalysisAnalyticsGraphLineBusinessBoardChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M3 3h2M3 5.5h4.5m4 0l-3 5l-3.5-2l-2 3"/></g>`), g.Group(children),
	)
}

func InterfaceContentEyeGlassesVisionSunglassesProtectionSpectaclesCorrectionSunEyeGlasses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9h-5v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1Zm8 0h-5v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1Zm-8 0h3m5 0V3a2 2 0 0 0-2-2h-1M.5 9V3a2 2 0 0 1 2-2h1"/>`), g.Group(children),
	)
}

func InterfaceContentFileEmptyCommonFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5Z"/><path d="M7.5.5v5h5"/></g>`), g.Group(children),
	)
}

func InterfaceContentFireLitFlameTorchTrending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.15.53a.39.39 0 0 0-.4 0a.26.26 0 0 0-.06.34C6.92 3 7.18 5.9 5.5 7.5a5.52 5.52 0 0 1-1.5-2A3.89 3.89 0 0 0 2 9a4.7 4.7 0 0 0 5 4.5c3.22 0 4.89-2 5-4.5c.13-3-2-6.69-5.85-8.47Z"/><path d="M9.5 9a2 2 0 0 1-2 2"/></g>`), g.Group(children),
	)
}

func InterfaceContentNotePadContentNotesBookNotepadNotebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 3.5v-3m3 3v-3m3 3v-3"/><rect width="13" height="11.5" x=".5" y="2" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceContentNotePadTextContentNotesBookNotepadNotebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 3.5v-3m3 3v-3m3 3v-3"/><rect width="13" height="11.5" x=".5" y="2" rx="1"/><path d="M3.5 7h7m-7 3h4"/></g>`), g.Group(children),
	)
}

func InterfaceContentSaveDiskFloppyElectronicsDeviceDiscComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 12.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-9l3-3h9a1 1 0 0 1 1 1Z"/><path d="M3.5 8.5h7v5h-7zm1-8h6v4h-6z"/></g>`), g.Group(children),
	)
}

func InterfaceCursorArrowOneMouseSelectCursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 10.5l-4-4l2-2a1 1 0 0 0-.5-1.68L1.74.53A1 1 0 0 0 .53 1.74L2.82 11a1 1 0 0 0 1.68.46l2-2l4 4Z"/>`), g.Group(children),
	)
}

func InterfaceCursorArrowTwoMouseSelectCursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.15 5.45a.5.5 0 0 0 0-1L1.83.56A1 1 0 0 0 .56 1.83L4.5 13.16a.5.5 0 0 0 1 0l2-5.66Z"/>`), g.Group(children),
	)
}

func InterfaceCursorHandHandSelectCursorFinger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.13a2 2 0 0 1 1.59 2.24l-.61 4.27a1 1 0 0 1-1 .86H4a1 1 0 0 1-.93-.63L2 10.21a2 2 0 0 1 1-2.53L4.35 7V2a1.5 1.5 0 0 1 3 0v3.5Z"/>`), g.Group(children),
	)
}

func InterfaceDashboardLayoutCircleAppApplicationDashboardHomeLayoutCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.25" cy="3.25" r="2.75"/><circle cx="10.75" cy="3.25" r="2.75"/><circle cx="3.25" cy="10.75" r="2.75"/><circle cx="10.75" cy="10.75" r="2.75"/></g>`), g.Group(children),
	)
}

func InterfaceDashboardLayoutOneCornersDashboardFrameLayoutMatOctagonSquareTriangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M.5 6.5L6 .5m7.5 6L8 .5m-7.5 7l5.5 6m7.5-6l-5.5 6"/></g>`), g.Group(children),
	)
}

func InterfaceDashboardLayoutSquareAppApplicationDashboardHomeLayoutSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="5" x=".5" y=".5" rx="1"/><rect width="5" height="5" x="8.5" y=".5" rx="1"/><rect width="5" height="5" x=".5" y="8.5" rx="1"/><rect width="5" height="5" x="8.5" y="8.5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceDashboardLayoutThreeAppApplicationDashboardHomeLayout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="7" x="8.5" y="6.5" rx=".5"/><rect width="5" height="3.01" x="8.5" y=".5" rx=".5"/><rect width="5" height="7" x=".5" y=".5" rx=".5"/><rect width="5" height="3.01" x=".5" y="10.49" rx=".5"/></g>`), g.Group(children),
	)
}

func InterfaceDashboardLayoutTwoCornersDashboardFrameLayoutCircleSquareCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5.5h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m7 13h2a1 1 0 0 0 1-1v-2m-13 0v2a1 1 0 0 0 1 1h2"/><circle cx="7" cy="7" r="3.5"/></g>`), g.Group(children),
	)
}

func InterfaceDeleteBinFiveRemoveDeleteEmptyBinTrashGarbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1 3.5h12m-10.5 0h9v9a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-9h0Zm2 0V3a2.5 2.5 0 0 1 5 0v.5"/>`), g.Group(children),
	)
}

func InterfaceDeleteBinOneRemoveDeleteEmptyBinTrashGarbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.5 5.5l-1 8h-7l-1-8M1 3.5h12m-8.54-.29V1.48a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v2"/>`), g.Group(children),
	)
}

func InterfaceDeleteBinThreeRemoveDeleteEmptyBinTrashGarbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12 .5l-1.4 12.11a1 1 0 0 1-1 .89H4.39a1 1 0 0 1-1-.89L2 .5m-1 0h12m-6 0v13m-4.54-9h9.08M2.98 9h8.04"/>`), g.Group(children),
	)
}

func InterfaceDeleteBinTwoRemoveDeleteEmptyBinTrashGarbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1 3.5h12m-10.5 0h9v9a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-9h0Zm2 0V3a2.5 2.5 0 0 1 5 0v.5m-4 2V11m3-5.5V11"/>`), g.Group(children),
	)
}

func InterfaceDeleteCircleButtonDeleteRemoveAddCircleButtons(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.12 4.88L4.88 9.12m0-4.24l4.24 4.24"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceDeleteOneRemoveAddButtonButtonsDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5.5l-13 13m0-13l13 13"/>`), g.Group(children),
	)
}

func InterfaceDeleteSquareButtonRemoveButtonsAddSquareDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.12 4.88L4.88 9.12m0-4.24l4.24 4.24"/><rect width="13" height="13" x=".5" y=".5" rx="3"/></g>`), g.Group(children),
	)
}

func InterfaceDeleteThreeRemoveCircleGarbageTrashDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m2.4 11.6l9.2-9.2m0 9.2L2.4 2.4"/></g>`), g.Group(children),
	)
}

func InterfaceDeleteTwoRemoveBoldAddButtonButtonsDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.19 3.05a1.06 1.06 0 0 0 0-1.49l-.75-.75a1.06 1.06 0 0 0-1.44 0L7 4.76L3.05.81a1.06 1.06 0 0 0-1.49 0l-.75.75a1.06 1.06 0 0 0 0 1.49l4 4l-4 3.95a1.06 1.06 0 0 0 0 1.49l.75.75a1.06 1.06 0 0 0 1.49 0l3.95-4l4 3.95a1.06 1.06 0 0 0 1.49 0l.75-.75a1.06 1.06 0 0 0 0-1.49L9.24 7Z"/>`), g.Group(children),
	)
}

func InterfaceDownloadBoxOneArrowBoxDownDownloadInternetNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 13.5h-1a1 1 0 0 1-1-1v-8h13v8a1 1 0 0 1-1 1h-1"/><path d="M4.5 11L7 13.5L9.5 11M7 13.5v-6M11.29 1a1 1 0 0 0-.84-.5h-6.9a1 1 0 0 0-.84.5L.5 4.5h13ZM7 .5v4"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadButtonOneArrowButtonDownDownloadInternetNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 4.75h1a.5.5 0 0 1 .5.5v7.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-7.5a.5.5 0 0 1 .5-.5h1m3.5-4v8"/><path d="m5 6.75l2 2l2-2"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadButtonTwoArrowBottomDownDownloadInternetNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 10.5v1a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M4 6l3 3.5L10 6M7 9.5v-9"/>`), g.Group(children),
	)
}

func InterfaceDownloadCircleArrowCircleDownDownloadInternetNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4 7l3 3.5L10 7m-3 3.5v-7"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadDesktopActionActionsComputerDesktopDeviceDisplayDownloadMonitorScreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 3H13a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-7A.5.5 0 0 1 1 3h1.5M7 11v2.5m-2 0h4M7 .5v6"/><path d="m5 4.5l2 2l2-2"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadLaptopArrowComputerDownDownloadInternetLaptopNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.91 9.5l-1.3 2.55a1 1 0 0 0 0 1a1 1 0 0 0 .87.47h11a1 1 0 0 0 .87-.47a1 1 0 0 0 0-1L12.09 9.5ZM5 4.5l2 2l2-2m-2 2v-6"/><path d="M2.5 4.63a1 1 0 0 0-.5.87v4h10v-4a1 1 0 0 0-.5-.87"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadSquareArrowDownDownloadInternetNetworkServerSquareUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m4 7l3 3.5L10 7m-3 3.5v-7"/></g>`), g.Group(children),
	)
}

func InterfaceDownloadWebsiteActionActionsComputerWebsiteDeviceDisplayDownloadMonitorScreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 7v6.5M4.5 11L7 13.5L9.5 11"/><path d="M2.5 13.5h-1a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-1m-11-10h13"/></g>`), g.Group(children),
	)
}

func InterfaceEditAttachmentOneAttachmentLinkPaperclipUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13 6.81l-5.95 6a2.48 2.48 0 0 1-3.54 0L1.73 11a2.53 2.53 0 0 1 0-3.55l6.34-6.36a2 2 0 0 1 2.84 0l.71.71a2 2 0 0 1 0 2.84L6 10.28a1 1 0 0 1-1.42 0l-.35-.36a1 1 0 0 1 0-1.42L8 4.76"/>`), g.Group(children),
	)
}

func InterfaceEditAttachmentTwoAttachmentLinkPaperclipUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.75 11.5V3A2.5 2.5 0 0 0 8.25.5h-2.5A2.5 2.5 0 0 0 3.25 3v8.5a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V4a1 1 0 0 0-1-1h-.5a1 1 0 0 0-1 1v5.5"/>`), g.Group(children),
	)
}

func InterfaceEditBinocularBinocularBinocularsViewZoom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3" cy="11" r="2.5"/><path d="M5.5 3V1.5a1 1 0 0 0-1-1H3.32a1 1 0 0 0-1 .81L.55 10.5m3.95-5h5M5.5 8v3"/><circle cx="11" cy="11" r="2.5"/><path d="M8.5 3V1.5a1 1 0 0 1 1-1h1.18a1 1 0 0 1 1 .81l1.79 9.19M8.5 8v3"/></g>`), g.Group(children),
	)
}

func InterfaceEditBombDeleteBombRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.5" cy="7.5" r="6"/><path d="m13.5.5l-2.76 2.76M3.5 7.5a3 3 0 0 1 3-3"/></g>`), g.Group(children),
	)
}

func InterfaceEditBrushFourDesignRubberStampSuppliesTool(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 10.5h11v2a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-2h0Zm9-4a1 1 0 0 1-1-1V3a2.5 2.5 0 0 0-5 0v2.5a1 1 0 0 1-1 1h-1a2 2 0 0 0-2 2v1a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2Z"/>`), g.Group(children),
	)
}

func InterfaceEditBrushOneBrushColorColorsDesignPaintPainting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 8.5h-9l-.76 3.8a1 1 0 0 0 .21.83a1 1 0 0 0 .77.37h8.56a1 1 0 0 0 .77-.37a1 1 0 0 0 .21-.83Zm0-3a1 1 0 0 1 1 1v2h-11v-2a1 1 0 0 1 1-1H4a1 1 0 0 0 1-1v-2a2 2 0 0 1 4 0v2a1 1 0 0 0 1 1Zm-3 8V11"/>`), g.Group(children),
	)
}

func InterfaceEditBrushThreeBrushColorColorsDesignPaintPaintingRollerRolling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="4" x="2.5" y=".5" rx="1"/><path d="M8 9.5v-1a2 2 0 0 0-2-2H2.5a2 2 0 0 1-2-2v-1a1 1 0 0 1 1-1h1m4 11v-3a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v3"/></g>`), g.Group(children),
	)
}

func InterfaceEditBrushTwoBrushColorColorsDesignPaintPainting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.13 12.06C4.6 13.6 2 14.11.5 12.57C2.5 10.5.5 9.5 2 8a2.9 2.9 0 1 1 4.09 4.1Z"/><path d="M12.92 1.08A2 2 0 0 0 11.44.5a2 2 0 0 0-1.44.67l-5.38 6A2.85 2.85 0 0 1 6.13 8a3 3 0 0 1 .77 1.31L12.83 4a2 2 0 0 0 .67-1.43a2 2 0 0 0-.58-1.49Z"/></g>`), g.Group(children),
	)
}

func InterfaceEditClipBinderClipClipperCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 6h1a1 1 0 0 1 1 1.13l-.73 5.5a1 1 0 0 1-1 .87H2.23a1 1 0 0 1-1-.87l-.72-5.5A1 1 0 0 1 1.5 6h1"/><path d="M10.46 13.5L8.33 4h.42a1.75 1.75 0 0 0 0-3.5h-3.5a1.75 1.75 0 0 0 0 3.5h.41l-2.12 9.5"/></g>`), g.Group(children),
	)
}

func InterfaceEditColorDropPickColorDropPickCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 9C12 5.5 7 .5 7 .5S2 5.5 2 9a4.77 4.77 0 0 0 5 4.5A4.77 4.77 0 0 0 12 9ZM7 .5v13m2.5-.6V3.44"/>`), g.Group(children),
	)
}

func InterfaceEditColorPaletteColorPaletteCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.5" cy="9" r="4"/><circle cx="9.5" cy="9" r="4"/><circle cx="7" cy="5" r="4"/></g>`), g.Group(children),
	)
}

func InterfaceEditColorTriangleColorTriangleCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M3.5 9h7L7 3.5L3.5 9z"/></g>`), g.Group(children),
	)
}

func InterfaceEditCopyClipboardCopyCutPaste(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 3.5v-1a1 1 0 0 0-1-1h-1m-2.5 9H1.5a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h1"/><rect width="7" height="8" x="6.5" y="5.5" rx="1"/><path d="M6.75.5h-4.5l.41 1.62a.49.49 0 0 0 .48.38h2.72a.49.49 0 0 0 .48-.38Zm1.75 8h3m-3 2h3"/></g>`), g.Group(children),
	)
}

func InterfaceEditCropArtboardCropDesignImagePicture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4.5h1.5a1 1 0 0 1 1 1V7"/><path d="M4.5.5v8a1 1 0 0 0 1 1h8m-9-5h-4m9 5v4"/></g>`), g.Group(children),
	)
}

func InterfaceEditCutCouponCutDiscountPricePricesScissors(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.49 10.5h2m2 0h2M2.19 4.93l5.8 3.33"/><circle cx="2.75" cy="2.75" r="2.25"/><path d="M2.19 9.07L13.5 2.55"/><circle cx="2.75" cy="11.25" r="2.25"/></g>`), g.Group(children),
	)
}

func InterfaceEditCutterCutterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 7.5l-2 2v4l4-4"/><path d="M11.67 1.33a2.82 2.82 0 0 0-4 0L2.5 6.5l4 4l5.17-5.17a2.82 2.82 0 0 0 0-4ZM6.5 6.5l3-3"/></g>`), g.Group(children),
	)
}

func InterfaceEditDesignToolSelectionWandSselectionWandObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l9-9m2-2l1-1M9 2V.5M12 5h1.5M11 7l1 1M6 2l1 1"/>`), g.Group(children),
	)
}

func InterfaceEditDrawingBoardBoardDesignDrawingEaselProcess(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 2.5h8a1 1 0 0 1 1 1V10h0H2h0V3.5a1 1 0 0 1 1-1ZM.5 10h13M7 2.5v-2M5.5 10L3 13.5M8.5 10l2.5 3.5"/>`), g.Group(children),
	)
}

func InterfaceEditExpandBigBiggerDesignExpandLargerResizeSizeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3m11 5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3"/><rect width="6" height="6" x="4" y="4" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceEditFlipBottomAlternateOneFlipBottomObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7a6.5 6.5 0 0 1-13 0M13 4.39a6.13 6.13 0 0 0-.76-1.3a6.34 6.34 0 0 0-1-1.09M1.05 4.39a6.13 6.13 0 0 1 .76-1.3A6.34 6.34 0 0 1 2.85 2M8.5.67a6.7 6.7 0 0 0-3 0"/>`), g.Group(children),
	)
}

func InterfaceEditFlipBottomFlipBottomObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7h13a6.5 6.5 0 0 1-13 0ZM13 4.39a6.13 6.13 0 0 0-.76-1.3a6.34 6.34 0 0 0-1-1.09M1.05 4.39a6.13 6.13 0 0 1 .76-1.3A6.34 6.34 0 0 1 2.85 2M8.5.67a6.7 6.7 0 0 0-3 0"/>`), g.Group(children),
	)
}

func InterfaceEditFlipDownDesignDownFlipReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m-3 0h-3m-5-2v1a1 1 0 0 0 1 1h1m-2-9v-3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v3m-13 4V7h13v1.5"/>`), g.Group(children),
	)
}

func InterfaceEditFlipHorizontalOneArrowDesignFlipReflect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 3l4 4l-4 4V3zm13 0l-4 4l4 4V3zM7 .5v1m0 2v1m0 2v1m0 2v1m0 2v1"/>`), g.Group(children),
	)
}

func InterfaceEditFlipHorizontalTwoArrowDesignFlipReflect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.5 3.75l2.93 3.07a.27.27 0 0 1 0 .36L1.5 10.25m11-6.5L9.57 6.82a.27.27 0 0 0 0 .36l2.93 3.07M7 .5v1m0 2v1m0 2v1m0 2v1m0 2v1"/>`), g.Group(children),
	)
}

func InterfaceEditFlipLeftAlternateOneFlipLeftObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5a6.5 6.5 0 0 0 0 13M9.61 1.05a6.13 6.13 0 0 1 1.3.76a6.34 6.34 0 0 1 1.09 1M9.61 13a6.13 6.13 0 0 0 1.3-.76a6.34 6.34 0 0 0 1.09-1m1.33-5.74a6.7 6.7 0 0 1 0 3"/>`), g.Group(children),
	)
}

func InterfaceEditFlipLeftDesignFlipReflectLeftHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h-1a1 1 0 0 0-1 1v1m0 3v3m2 5h-1a1 1 0 0 1-1-1v-1m9 2h3a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-3m-4 13H7V.5H5.5"/>`), g.Group(children),
	)
}

func InterfaceEditFlipLeftFlipLeftObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V.5a6.5 6.5 0 0 0 0 13ZM9.61 1.05a6.13 6.13 0 0 1 1.3.76a6.34 6.34 0 0 1 1.09 1M9.61 13a6.13 6.13 0 0 0 1.3-.76a6.34 6.34 0 0 0 1.09-1m1.33-5.74a6.7 6.7 0 0 1 0 3"/>`), g.Group(children),
	)
}

func InterfaceEditFlipRightAlternateOneFlipRightObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5a6.5 6.5 0 0 1 0 13M4.39 1.05a6.13 6.13 0 0 0-1.3.76A6.34 6.34 0 0 0 2 2.85M4.39 13a6.13 6.13 0 0 1-1.3-.76a6.34 6.34 0 0 1-1.09-1M.67 5.5a6.7 6.7 0 0 0 0 3"/>`), g.Group(children),
	)
}

func InterfaceEditFlipRightDesignFlipReflectRightHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m0 3v3m-2 5h1a1 1 0 0 0 1-1v-1m-9 2h-3a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h3m4 13H7V.5h1.5"/>`), g.Group(children),
	)
}

func InterfaceEditFlipRightFlipRightObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V.5a6.5 6.5 0 0 1 0 13ZM4.39 1.05a6.13 6.13 0 0 0-1.3.76A6.34 6.34 0 0 0 2 2.85M4.39 13a6.13 6.13 0 0 1-1.3-.76a6.34 6.34 0 0 1-1.09-1M.67 5.5a6.7 6.7 0 0 0 0 3"/>`), g.Group(children),
	)
}

func InterfaceEditFlipTopAlternateOneFlipTopObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7a6.5 6.5 0 0 1 13 0M1.05 9.61a6.13 6.13 0 0 0 .76 1.3a6.34 6.34 0 0 0 1 1.09M13 9.61a6.13 6.13 0 0 1-.76 1.3a6.34 6.34 0 0 1-1 1.09M5.5 13.33a6.7 6.7 0 0 0 3 0"/>`), g.Group(children),
	)
}

func InterfaceEditFlipTopFlipTopObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7H.5a6.5 6.5 0 0 1 13 0ZM1.05 9.61a6.13 6.13 0 0 0 .76 1.3a6.34 6.34 0 0 0 1 1.09M13 9.61a6.13 6.13 0 0 1-.76 1.3a6.34 6.34 0 0 1-1 1.09M5.5 13.33a6.7 6.7 0 0 0 3 0"/>`), g.Group(children),
	)
}

func InterfaceEditFlipUpDesignUpFlipReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2.5v-1a1 1 0 0 0-1-1h-1m-3 0h-3m-5 2v-1a1 1 0 0 1 1-1h1m-2 9v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-3m-13-4V7h13V5.5"/>`), g.Group(children),
	)
}

func InterfaceEditFlipVerticalOneArrowDesignFlipReflectUpDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25 1.5L7.18 4.43a.27.27 0 0 1-.36 0L3.75 1.5m6.5 11L7.18 9.57a.27.27 0 0 0-.36 0L3.75 12.5M13.5 7h-1m-2 0h-1m-2 0h-1m-2 0h-1m-2 0h-1"/>`), g.Group(children),
	)
}

func InterfaceEditFlipVerticalTwoArrowDesignFlipReflectUpDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11 13.5l-4-4l-4 4h8zm0-13l-4 4l-4-4h8zM13.5 7h-1m-2 0h-1m-2 0h-1m-2 0h-1m-2 0h-1"/>`), g.Group(children),
	)
}

func InterfaceEditGlueGlueCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 8v4.5a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1V8A2.5 2.5 0 0 1 5 5.5h4A2.5 2.5 0 0 1 11.5 8ZM7.5.5h-1L5 5.5h4L7.5.5zM5 11h4M2.5 8.5h9"/>`), g.Group(children),
	)
}

func InterfaceEditGridGridLayoutLayoutsModule(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M5 .5v13m4-13v13M13.5 5H.5m13 4H.5"/></g>`), g.Group(children),
	)
}

func InterfaceEditGridOffGridLayoutLayoutsModuleOffCloseDenySlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13M1.79 1.79a1 1 0 0 1 .71-.29h9a1 1 0 0 1 1 1v9a1 1 0 0 1-.29.71M9 12.5H2.5a1 1 0 0 1-1-1V5m3.75-3.5v3.75m3.5-3.75v7.25m3.75-3.5H5.25m7.25 3.5H8.75m-3.5-.25v4m3.5-.5v.5M5.5 8.75h-4m.5-3.5h-.5"/>`), g.Group(children),
	)
}

func InterfaceEditLayerAddOneLayerAddDesignPlusLayersSquareBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.19v4.62M3.19 5.5h4.62"/><rect width="10" height="10" x=".5" y=".5" rx="1"/><path d="M3.5 13.5h9a1 1 0 0 0 1-1v-9"/></g>`), g.Group(children),
	)
}

func InterfaceEditLayerAddTwoLayerAddDesignPlusLayersSquareBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x=".5" y="3.5" rx="1"/><path d="M3.5.5h9a1 1 0 0 1 1 1v9M5.5 6v5M8 8.5H3"/></g>`), g.Group(children),
	)
}

func InterfaceEditMagicWandDesignMagicStarSuppliesToolWand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.64 1.87l-.84 2.48a.41.41 0 0 0 0 .37l1.57 2.1a.4.4 0 0 1-.33.64h-2.62a.43.43 0 0 0-.33.17l-1.46 2.1a.4.4 0 0 1-.71-.11l-.78-2.5a.38.38 0 0 0-.26-.26l-2.5-.78a.4.4 0 0 1-.11-.71l2.14-1.51a.43.43 0 0 0 .17-.33V.91a.4.4 0 0 1 .6-.33l2.1 1.57a.41.41 0 0 0 .37.05l2.48-.84a.4.4 0 0 1 .51.51Zm-5.6 5.09L.5 13.5"/>`), g.Group(children),
	)
}

func InterfaceEditMagnetDesignMagnetSnapSuppliesToTool(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.81 12.07a4.86 4.86 0 0 1-6.88-6.88L6.62.5l2.19 2.19L4.51 7A1.77 1.77 0 0 0 7 9.49l4.3-4.3l2.2 2.19Zm.38-4.76l2.19 2.19M4.5 2.62l2.19 2.19"/>`), g.Group(children),
	)
}

func InterfaceEditPaintColorColorsDesignPaintPaintingPalette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8.5" cy="4" r="1"/><circle cx="4.5" cy="9.5" r=".5"/><circle cx="4.5" cy="5.5" r="1"/><path d="M9.52 12.28a1 1 0 0 0-.65-.88a2 2 0 0 1 .63-3.9h1.87a2 2 0 0 0 1.89-2.67a6.5 6.5 0 1 0-6.13 8.67a6.3 6.3 0 0 0 1.74-.24a.9.9 0 0 0 .65-.98Z"/></g>`), g.Group(children),
	)
}

func InterfaceEditPathfinderDividePathfinderDivideWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5a1 1 0 0 0 1 1m0-9a1 1 0 0 0-1 1m9 0a1 1 0 0 0-1-1M4 .5h2M.5 4v2m9-1.5h3a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-3h4a1 1 0 0 0 1-1Z"/>`), g.Group(children),
	)
}

func InterfaceEditPathfinderIntersectPathfinderIntersectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5a1 1 0 0 0 1 1m0-9a1 1 0 0 0-1 1m9 0a1 1 0 0 0-1-1M4 .5h2M.5 4v2m4 6.5a1 1 0 0 0 1 1m8-8a1 1 0 0 0-1-1m0 9a1 1 0 0 0 1-1m-5.5 1h2M13.5 8v2m-4-1.5v-4h-4a1 1 0 0 0-1 1v4h4a1 1 0 0 0 1-1Z"/>`), g.Group(children),
	)
}

func InterfaceEditPathfinderMergePathfinderMergeWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 9.5h1a1 1 0 0 0 1-1v-1m0-3v-3a1 1 0 0 0-1-1h-7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h3m0 3a1 1 0 0 0 1 1m8-8a1 1 0 0 0-1-1m0 9a1 1 0 0 0 1-1m-4-8h.5m-2 9h2M13.5 8v2m-9-.5v.5"/>`), g.Group(children),
	)
}

func InterfaceEditPathfinderOutlinePathfinderOutlineWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 9.5a1 1 0 0 0 1-1m-5 1H6m3.5-5V6m-4-1.5a1 1 0 0 0-1 1"/><path d="M.5 1.5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v3h3a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-3h-3a1 1 0 0 1-1-1Zm7.5 3h1.5M4.5 8v1.5"/></g>`), g.Group(children),
	)
}

func InterfaceEditPenOneContentCreationEditPenWrite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 13.5l2-2m2.21.79a1 1 0 0 1-1.42 0l-1.58-1.58a1 1 0 0 1 0-1.42l8.17-8.17a2.12 2.12 0 0 1 3 3Z"/><path d="M11.5 5.5L7.21 1.21a1 1 0 0 0-1.42 0L2.5 4.5"/></g>`), g.Group(children),
	)
}

func InterfaceEditPenThreeContentCreationEditPenPensWrite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3a2.5 2.5 0 0 0-5 0v7.5l2.5 3l2.5-3Zm-5 1.5h5m2-2a2 2 0 0 1 4 0v8a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1Zm2 9v2"/><path d="M7.5 4.5h5a1 1 0 0 1 1 1v4"/></g>`), g.Group(children),
	)
}

func InterfaceEditPenTwoContentCreationEditFountainPenWrite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 12.5a2.72 2.72 0 0 0 0-4a2.72 2.72 0 0 0-4 0c-1 1-1 5-1 5s4 0 5-1Z"/><path d="M12.92 1.08a2 2 0 0 0-2.64-.15L4.5 5l-.71 2.64a2.87 2.87 0 0 1 1.71.86a2.87 2.87 0 0 1 .86 1.71L9 9.5l4.07-5.78a2 2 0 0 0-.15-2.64ZM.5 13.5l3.25-3.25"/></g>`), g.Group(children),
	)
}

func InterfaceEditPencilChangeEditModifyPencilWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 12.24L.5 13.5L1.76 9L10 .8a1 1 0 0 1 1.43 0l1.77 1.78a1 1 0 0 1 0 1.42Z"/>`), g.Group(children),
	)
}

func InterfaceEditPickerColorColorsDesignDropperEyeEyedropEyedropperPaintingPicker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.46 1.54a3.52 3.52 0 0 0-5 0L2.87 6.13a3 3 0 0 0-.53 3.53L.79 11.21a1 1 0 0 0 0 1.41l.59.59a1 1 0 0 0 1.41 0l1.55-1.55a3 3 0 0 0 3.53-.53l4.59-4.59a3.52 3.52 0 0 0 0-5ZM4.5 1.5l8 8"/>`), g.Group(children),
	)
}

func InterfaceEditPinOnePinPushThumbtack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="5" r="4.5"/><path d="m.5 13.5l5.58-5.07"/><circle cx="8.5" cy="3.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceEditPinThreePinPushThumbtack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="5" r="4.5"/><path d="M10.25 3.67a1.5 1.5 0 0 0-2.31-.23M.5 13.5l5.58-5.07"/></g>`), g.Group(children),
	)
}

func InterfaceEditPinTwoPinPushThumbtack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.73 7.65L13 5.54A1 1 0 0 0 13.21 4L10 .79A1 1 0 0 0 8.46 1L6.3 4.23l-4.49 1a.6.6 0 0 0-.29 1l6.15 6.16a.61.61 0 0 0 1-.3ZM4.59 9.38L.5 13.5"/>`), g.Group(children),
	)
}

func InterfaceEditPrinterPrinterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 8.5h9a1 1 0 0 1 1 1v4h0h-11h0v-4a1 1 0 0 1 1-1Z"/><path d="M1.5 9.5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1M3 .5h8v5H3zM4 11h6"/></g>`), g.Group(children),
	)
}

func InterfaceEditQuillChangeEditFeatherModifyQuillWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.51 10.05a10.21 10.21 0 0 0 3 .45A10 10 0 0 0 13.5 1a10.16 10.16 0 0 0-3-.45a10 10 0 0 0-9.99 9.5Z"/><path d="M1 13.5a9.65 9.65 0 0 1-.49-3.45M8.01 5l2.57 2.57"/></g>`), g.Group(children),
	)
}

func InterfaceEditRotateAngleRotateAngleCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13h13"/><path d="M7.5 13.5a7 7 0 0 0-7-7m0 7L3 11m2.5-2.5L7 7m2-2l1.5-1.5m2-2l1-1"/></g>`), g.Group(children),
	)
}

func InterfaceEditRulerRulerCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.502 9.496L9.503.495l4.002 4.002l-9.001 9.001zM7.5 2.5L9 4M5 5l1.5 1.5m-4 1L4 9"/>`), g.Group(children),
	)
}

func InterfaceEditScissorsClipboardCopyCutPasteRightScissors(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.19 4.93l11.31 6.52"/><circle cx="2.75" cy="2.75" r="2.25"/><path d="M2.19 9.07L13.5 2.55"/><circle cx="2.75" cy="11.25" r="2.25"/></g>`), g.Group(children),
	)
}

func InterfaceEditSelectAreaCircleDashSelectAreaObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5.67a6.7 6.7 0 0 0-3 0m-1.91.8a7 7 0 0 0-1.19.93a7 7 0 0 0-.93 1.19M.67 5.5a6.7 6.7 0 0 0 0 3m.8 1.91a7 7 0 0 0 .93 1.19a7 7 0 0 0 1.19.93m1.91.8a6.7 6.7 0 0 0 3 0m1.91-.8a7 7 0 0 0 1.19-.93a7 7 0 0 0 .93-1.19m.8-1.91a6.7 6.7 0 0 0 0-3m-.8-1.91a7 7 0 0 0-.93-1.19a7 7 0 0 0-1.19-.93"/>`), g.Group(children),
	)
}

func InterfaceEditSelectAreaRectangleDashSelectAreaObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3m11 5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3"/>`), g.Group(children),
	)
}

func InterfaceEditSelectBackBehindBoxDesignLayerLayersSelectSendToBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 9.5a1 1 0 0 0 1 1m0-10a1 1 0 0 0-1 1m10 0a1 1 0 0 0-1-1m0 10a1 1 0 0 0 1-1"/><path d="M10.5 13.5h-8a2 2 0 0 1-2-2v-8m7-3h2m-2 10h2m4-6v2m-10-2v2"/></g>`), g.Group(children),
	)
}

func InterfaceEditSelectFrameCursorFrameSelect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5.5h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m7 13h2a1 1 0 0 0 1-1v-2m-13 0v2a1 1 0 0 0 1 1h2"/>`), g.Group(children),
	)
}

func InterfaceEditSelectFrontDesignFrontLayerLayersSelectSendToTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x="3.5" y=".5" rx="1"/><path d="M.5 12.5a1 1 0 0 0 1 1m3 0H6m3 0h1.5M.5 8v1.5m0-6V5"/></g>`), g.Group(children),
	)
}

func InterfaceEditSkullOneCrashDeathDeleteDieErrorGarbageRemoveSkullTrash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 6.5a6 6 0 1 0-9.5 4.87v1.13a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-1.13A6 6 0 0 0 13 6.5Z"/><circle cx="4.5" cy="7" r=".5"/><circle cx="9.5" cy="7" r=".5"/><path d="M6 11.5v2m2-2v2"/></g>`), g.Group(children),
	)
}

func InterfaceEditSprayCanColorColorsDesignPaintPaintingSpray(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6" height="9" x=".5" y="4.5" rx="1"/><path d="M3.5 2v2.5m5-3l5-1m-5 3l5 1M2.5 2h2"/></g>`), g.Group(children),
	)
}

func InterfaceEditSwatchColorColorsDesignPaintingPaletteSampleSwatch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5.5h3a1 1 0 0 1 1 1V11A2.5 2.5 0 0 1 3 13.5h0A2.5 2.5 0 0 1 .5 11V1.5a1 1 0 0 1 1-1Z"/><path d="M5.5 5L9 1.48a1 1 0 0 1 1.41 0l2.11 2.12a1 1 0 0 1 0 1.41l-7.75 7.76"/><path d="M9 8.5h3.5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H3m-2.5-9h5m-5 4h5"/></g>`), g.Group(children),
	)
}

func InterfaceEditTypewriterTypewriterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 3.5h13v3H.5zM3 .5h8a.5.5 0 0 1 .5.5v2.5h0h-9h0V1A.5.5 0 0 1 3 .5Z"/><circle cx="3" cy="8" r="1.5"/><circle cx="11" cy="8" r="1.5"/><path d="M2.5 9.41v3.09a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V9.41m-9 .09h9"/></g>`), g.Group(children),
	)
}

func InterfaceEditViewEyeEyeballOpenView(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.23 6.33a1 1 0 0 1 0 1.34C12.18 8.8 9.79 11 7 11S1.82 8.8.77 7.67a1 1 0 0 1 0-1.34C1.82 5.2 4.21 3 7 3s5.18 2.2 6.23 3.33Z"/><circle cx="7" cy="7" r="2"/></g>`), g.Group(children),
	)
}

func InterfaceEditViewOffDisableEyeEyeballHideOffView(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.29 5.4c.38.34.7.67.94.93a1 1 0 0 1 0 1.34C12.18 8.8 9.79 11 7 11h-.4m-2.73-.87a12.4 12.4 0 0 1-3.1-2.46a1 1 0 0 1 0-1.34C1.82 5.2 4.21 3 7 3a6.56 6.56 0 0 1 3.13.87M12.5 1.5l-11 11"/><path d="M5.59 8.41A2 2 0 0 1 5 7a2 2 0 0 1 2-2a2 2 0 0 1 1.41.59M8.74 8a2 2 0 0 1-.74.73"/></g>`), g.Group(children),
	)
}

func InterfaceEditWriteCircleChangeCircleEditModifyPencilWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7A6.5 6.5 0 1 1 7 .5"/><path d="m10.5.5l-5 5l-1 4l4-1l5-5"/></g>`), g.Group(children),
	)
}

func InterfaceEditWriteOneEditEditionFormPenTextWrite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5h11m-5-3.5l-3 .54L4 7.5L10.73.79a1 1 0 0 1 1.42 0l1.06 1.06a1 1 0 0 1 0 1.42Z"/>`), g.Group(children),
	)
}

func InterfaceEditWriteTwoChangeDocumentEditModifyPaperPencilWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 9l-3 .54L5 6.5L10.73.79a1 1 0 0 1 1.42 0l1.06 1.06a1 1 0 0 1 0 1.42Z"/><path d="M12 9.5v3a1 1 0 0 1-1 1H1.5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h3"/></g>`), g.Group(children),
	)
}

func InterfaceEditZoomInEnhanceGlassInMagnifyMagnifyingZoom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.92" cy="5.92" r="5.42"/><path d="M13.5 13.5L9.75 9.75M6 3.5v5M3.5 6h5"/></g>`), g.Group(children),
	)
}

func InterfaceEditZoomOutGlassMagnifyingOutReduceZoom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.92" cy="5.92" r="5.42"/><path d="M13.5 13.5L9.75 9.75M3.5 6h5"/></g>`), g.Group(children),
	)
}

func InterfaceFavoriteAwardRibbonRewardLikeSocialRatingMedia(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.87" cy="5" r="4.5"/><circle cx="6.87" cy="5" r="2"/><path d="m6 9.42l-.88 3.7a.51.51 0 0 1-.26.33a.54.54 0 0 1-.43 0L1.11 12a.51.51 0 0 1-.18-.78L3.5 8M8 9.37l.9 3.75a.5.5 0 0 0 .27.33a.51.51 0 0 0 .42 0l3.3-1.45a.5.5 0 0 0 .28-.35a.48.48 0 0 0-.1-.43l-2.68-3.41"/></g>`), g.Group(children),
	)
}

func InterfaceFavoriteDislikeOneRewardDownThumbHandSocialMediaDislikeRating(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.37 8.15l2.54 4.06a1.09 1.09 0 0 0 .94.52h0A1.11 1.11 0 0 0 8 11.63V8.72h4.39a1.15 1.15 0 0 0 1.1-1.32l-.8-5.16a1.14 1.14 0 0 0-1.13-1H5a2 2 0 0 0-.9.21l-.72.36m-.01 6.34V1.84M1 1.84h2.37v6.31h0H1a.5.5 0 0 1-.5-.5V2.34a.5.5 0 0 1 .5-.5Z"/>`), g.Group(children),
	)
}

func InterfaceFavoriteGiveHeartRewardSocialRatingMediaHeartHand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5h2A2.5 2.5 0 0 1 5 11h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5h-8m8.25-5l-4-3.7c-2.18-2.19 1-6.43 4-3c3-3.42 6.21.82 4 3Z"/>`), g.Group(children),
	)
}

func InterfaceFavoriteHeartRewardSocialRatingMediaHeartItLikeFavoriteLove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7 12.45l-5.52-5c-3-3 1.41-8.76 5.52-4.1c4.11-4.66 8.5 1.12 5.52 4.1Z"/>`), g.Group(children),
	)
}

func InterfaceFavoriteLikeOneRewardSocialUpRatingMediaLikeThumbHand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.37 5.85l2.54-4.06a1.09 1.09 0 0 1 .94-.52h0A1.11 1.11 0 0 1 8 2.37v2.91h4.39a1.15 1.15 0 0 1 1.1 1.32l-.8 5.16a1.14 1.14 0 0 1-1.13 1H5a2 2 0 0 1-.9-.21l-.72-.36m-.01-6.34v6.31M1 5.85h2.37v6.31h0H1a.5.5 0 0 1-.5-.5V6.35a.5.5 0 0 1 .5-.5Z"/>`), g.Group(children),
	)
}

func InterfaceFavoriteStarRewardRatingRateSocialStarMediaFavoriteLikeStars(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.49 1.09L9.08 4.3a.51.51 0 0 0 .41.3l3.51.52a.54.54 0 0 1 .3.93l-2.53 2.51a.53.53 0 0 0-.16.48l.61 3.53a.55.55 0 0 1-.8.58l-3.16-1.67a.59.59 0 0 0-.52 0l-3.16 1.67a.55.55 0 0 1-.8-.58L3.39 9a.53.53 0 0 0-.16-.48L.67 6.05A.54.54 0 0 1 1 5.12l3.51-.52a.51.51 0 0 0 .41-.3l1.59-3.21a.54.54 0 0 1 .98 0Z"/>`), g.Group(children),
	)
}

func InterfaceFileAddAlternateFileCommonAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 5.5v-4a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1h-5"/><path d="M8.5.5v5h5m-10 2v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func InterfaceFileAddFileCommonAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5Zm-6-6.75v3.5M4.75 7.5h3.5"/>`), g.Group(children),
	)
}

func InterfaceFileBookmarkTextCommonBookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1H9L12.5 4Z"/><path d="m8 7.5l-2-2l-2 2v-7h4v7z"/></g>`), g.Group(children),
	)
}

func InterfaceFileCheckFileCommonCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5Z"/><path d="m4.5 8.5l1.5 1l2.5-4"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardAddEditTaskEditionAddClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="M7 6v4m2-2H5"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardBlockEditTaskEditionBlockClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 13.5h-4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1H3m5 0h1.5a1 1 0 0 1 1 1v2"/><rect width="5" height="2.5" x="3" y=".5" rx="1"/><circle cx="10.25" cy="10.25" r="3.25"/><path d="m7.95 12.55l4.6-4.6"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardCheckCheckmarkEditTaskEditionChecklistCheckSuccessClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="m4.5 8.5l2 1.5L9 6"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardDeleteEditTaskEditionDeleteClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="m5.5 6.5l3 3m0-3l-3 3"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardRemoveEditTaskEditionRemoveDeleteClipboardForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="M9 8H5"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardTextEditionFormTaskChecklistEditClipboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/><path d="M4.5 5.5h5M4.5 8h5m-5 2.5h5"/></g>`), g.Group(children),
	)
}

func InterfaceFileClipboardWorkPlainClipboardTaskListCompanyOffice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 1.5H11a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h1.5"/><rect width="5" height="2.5" x="4.5" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceFileDeleteAlternateFileCommonDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 7V1.5a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1H8"/><path d="M8.5.5v5h5M4.74 9.26L.5 13.5m0-4.24l4.24 4.24"/></g>`), g.Group(children),
	)
}

func InterfaceFileDeleteFileCommonDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5ZM8.74 6L4.5 10.24M4.5 6l4.24 4.24"/>`), g.Group(children),
	)
}

func InterfaceFileDoubleFileCommonDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V1.5a1 1 0 0 1 1-1h4.5l3 3Z"/><path d="M9.5 13.5h-7a1 1 0 0 1-1-1v-9"/></g>`), g.Group(children),
	)
}

func InterfaceFileFolderWorkOfficeCompanyFolderSuppliesFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.91 12.56l-.41-7A.5.5 0 0 1 1 5h4.61a.51.51 0 0 1 .49.38L6.5 7H13a.5.5 0 0 1 .5.54l-.39 5a1 1 0 0 1-1 .92H1.91a1 1 0 0 1-1-.9ZM3.5 3V1A.5.5 0 0 1 4 .5h8.5a.5.5 0 0 1 .5.5v4M7.5 3h3"/>`), g.Group(children),
	)
}

func InterfaceFileMultipleDoubleCommonFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="11" x="2" y="2.5" rx="1"/><path d="M4 5h4M4 7.5h4M4 10h2M4.5.5H11a1 1 0 0 1 1 1V11"/></g>`), g.Group(children),
	)
}

func InterfaceFileRemoveAlternateFileCommonRemoveMinusSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 7.5v-6a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1h-6"/><path d="M8.5.5v5h5m-13 5h5"/></g>`), g.Group(children),
	)
}

func InterfaceFileRemoveFileCommonRemoveMinusSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5Zm-7.75-5h3.5"/>`), g.Group(children),
	)
}

func InterfaceFileSettingFileCommonSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.39 4V1.5a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1h-3.5"/><path d="M8.39.5v5h5m-9.75 1V8m-3.03.25l1.3.75m-1.3 2.75l1.3-.75m1.73 2.5V12m3.03-.25L5.37 11m1.3-2.75L5.37 9"/><circle cx="3.64" cy="10" r="2"/></g>`), g.Group(children),
	)
}

func InterfaceFileStickyNoteEmptyCommonFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 13.5h-7a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7Z"/><path d="M8.5 9v4.5l5-5H9a.5.5 0 0 0-.5.5Z"/></g>`), g.Group(children),
	)
}

func InterfaceFileTextTextCommonFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h5l5 5Zm-8-8h2m-2 3h5m-5 3h5"/>`), g.Group(children),
	)
}

func InterfaceFileZipFileCommonZip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h6l4 4ZM4 3.5h3m-3 3h3m-3 3h3M5.5.5v10"/>`), g.Group(children),
	)
}

func InterfaceFolderAddAddFolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 11.75v-9.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.31 1.24h6a1 1 0 0 1 1 1v7.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1Zm6.63-5.5v3.5M5.38 8h3.5"/>`), g.Group(children),
	)
}

func InterfaceFolderBlockBlockFolderSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.75 12H1.5a1 1 0 0 1-1-1V1.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.28 1.24h6a1 1 0 0 1 1 1v1.75"/><circle cx="10.25" cy="10.25" r="3.25"/><path d="m7.95 12.55l4.6-4.6"/></g>`), g.Group(children),
	)
}

func InterfaceFolderCheckRemoveCheckFolderSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 11.75v-9.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.31 1.24h6a1 1 0 0 1 1 1v7.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1Z"/><path d="m5 8.75l1.5 1l2.5-4"/></g>`), g.Group(children),
	)
}

func InterfaceFolderDeleteRemoveMinusFolderSubtractDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 11.75v-9.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.31 1.24h6a1 1 0 0 1 1 1v7.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1ZM5 8h3.5"/>`), g.Group(children),
	)
}

func InterfaceFolderEmptyFolder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6A1.5 1.5 0 0 0 12 4.5H7l-1.44-3H2A1.5 1.5 0 0 0 .5 3v8A1.5 1.5 0 0 0 2 12.5h10a1.5 1.5 0 0 0 1.5-1.5Z"/>`), g.Group(children),
	)
}

func InterfaceFolderRemoveRemoveMinusFolderSubtractDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 11.75v-9.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.31 1.24h6a1 1 0 0 1 1 1v7.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1Zm8-5l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func InterfaceFolderZipZipFolderCompact(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 6.25h3m-3 3h3m-1.5-6v7m-9 1.5v-9.5a1 1 0 0 1 1-1h3.69a1 1 0 0 1 1 .76l.31 1.24h6a1 1 0 0 1 1 1v7.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1Z"/>`), g.Group(children),
	)
}

func InterfaceGeometricCircleGeometricCircleRoundDesignShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="7" cy="7" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`), g.Group(children),
	)
}

func InterfaceGeometricPentagonPentagonDesignGeometricShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6L7 .5L.5 6l3 7.5h7l3-7.5z"/>`), g.Group(children),
	)
}

func InterfaceGeometricPolygonPolygonOctangleDesignGeometricShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 13.5h-5l-4-4v-5l4-4h5l4 4v5l-4 4z"/>`), g.Group(children),
	)
}

func InterfaceGeometricSquareSquareGeometricDesignShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="13" height="13" x=".5" y=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`), g.Group(children),
	)
}

func InterfaceGeometricTriangleGeometricTriangleShapeDesignShapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.46 2a.55.55 0 0 0-.92 0l-6 9.5a.5.5 0 0 0 0 .5a.54.54 0 0 0 .46.25h12a.54.54 0 0 0 .46-.25a.5.5 0 0 0 0-.5Z"/>`), g.Group(children),
	)
}

func InterfaceHelpCustomerSupportFiveCustomerHeadsetHelpPhoneSupport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3" height="6" x="2.5" y="7.5" rx="1"/><rect width="3" height="6" x="8.5" y="7.5" rx="1"/><path d="M.5 9.5V7a6.5 6.5 0 0 1 13 0v2.5"/></g>`), g.Group(children),
	)
}

func InterfaceHelpCustomerSupportFourCustomerHeadsetHelpPhoneSupport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 7.5h1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V7a6.5 6.5 0 0 1 13 0v5.5a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h1"/>`), g.Group(children),
	)
}

func InterfaceHelpCustomerSupportHumanOneCustomerHeadphonesHelpHeadsetPersonProfileSuuport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="M7 4.5v-4m2.5 13V12h2a1 1 0 0 0 1-1V7a5.5 5.5 0 0 0-3-4.9m-5 0a5.5 5.5 0 0 0 0 9.8v1.6"/></g>`), g.Group(children),
	)
}

func InterfaceHelpCustomerSupportOneCustomerHeadsetHelpMicrophonePhoneSupport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 7V4.37A3.93 3.93 0 0 1 7 .5h0a3.93 3.93 0 0 1 4 3.87V7M1.5 5.5h1A.5.5 0 0 1 3 6v3a.5.5 0 0 1-.5.5h-1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Zm11 4h-1A.5.5 0 0 1 11 9V6a.5.5 0 0 1 .5-.5h1a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1ZM9 12.25a2 2 0 0 0 2-2h0V8m-2 4.25a1.25 1.25 0 0 1-1.25 1.25h-1.5a1.25 1.25 0 0 1 0-2.5h1.5A1.25 1.25 0 0 1 9 12.25Z"/>`), g.Group(children),
	)
}

func InterfaceHelpCustomerSupportTwoCustomerHeadphonesHeadsetHelpMicrophonePhonePersonSupport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 6h1a.5.5 0 0 1 .5.5V9a.5.5 0 0 1-.5.5h-1a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1Zm11 3.5h-1A.5.5 0 0 1 11 9V6.5a.5.5 0 0 1 .5-.5h1a1 1 0 0 1 1 1v1.5a1 1 0 0 1-1 1Zm-3 2.75a2 2 0 0 0 2-2h0V9.5"/><path d="M8.25 11a1.25 1.25 0 0 1 0 2.5h-1.5a1.25 1.25 0 0 1 0-2.5ZM2.5 6V5a4.5 4.5 0 0 1 9 0v1m-6-2v1.5m3-1.5v1.5m-3 2c0 1.33 3 1.33 3 0"/></g>`), g.Group(children),
	)
}

func InterfaceHelpQuestionCircleCircleFaqFrameHelpInfoMarkMoreQueryQuestion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="7" cy="7" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5A1.5 1.5 0 1 1 7 7v1"/><path fill="currentColor" d="M7 9.5a.75.75 0 1 0 .75.75A.76.76 0 0 0 7 9.5Z"/>`), g.Group(children),
	)
}

func InterfaceHelpQuestionMessageBubbleHelpMarkMessageQueryQuestionSpeech(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5.5h-11a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h2v2.5L6.62 11h5.88a1 1 0 0 0 1-1V1.5a1 1 0 0 0-1-1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 4.5A1.5 1.5 0 1 1 7 6v.5"/><path fill="currentColor" d="M7 8a.75.75 0 1 0 .75.75A.76.76 0 0 0 7 8Z"/>`), g.Group(children),
	)
}

func InterfaceHelpQuestionSquareFrameHelpMarkQueryQuestionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="13" height="13" x=".5" y=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5A1.5 1.5 0 1 1 7 7v1"/><path fill="currentColor" d="M7 9.5a.75.75 0 1 0 .75.75A.76.76 0 0 0 7 9.5Z"/>`), g.Group(children),
	)
}

func InterfaceHierarchyOneNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="7" r="2"/><circle cx="11.5" cy="2.5" r="2"/><circle cx="11.5" cy="11.5" r="2"/><path d="m3.71 5.41l5.85-2.43M3.71 8.59l5.85 2.43"/></g>`), g.Group(children),
	)
}

func InterfaceHierarchyThreeNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3.5" height="3.5" x=".5" y="10" rx=".5"/><rect width="3.5" height="3.5" x="10" y="10" rx=".5"/><rect width="4" height="4" x="5" y=".5" rx=".5"/><path d="M4 12h6M5.09 4.29L2.5 10m6.41-5.71L11.5 10"/></g>`), g.Group(children),
	)
}

func InterfaceHierarchyTwoNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="1.5"/><circle cx="2" cy="11.5" r="1.5"/><circle cx="7" cy="11.5" r="1.5"/><circle cx="12" cy="11.5" r="1.5"/><path d="M2 10V8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2M7 4v6"/></g>`), g.Group(children),
	)
}

func InterfaceHomeFiveDoorEntranceHomeHouseMapRoofRoundWindow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 8.5L7 2l6.5 6.5"/><path d="M2.5 6.5v7h9v-7m-4.5 7v-3"/><circle cx="7" cy="6.75" r="1.25"/></g>`), g.Group(children),
	)
}

func InterfaceHomeFourDoorEntrnaceMap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="13" x="2.5" y=".5" rx="1"/><path d="M8 7.5h1"/></g>`), g.Group(children),
	)
}

func InterfaceHomeOneHomeHouseMapRoof(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.94a1 1 0 0 0-.32-.74L7 .5L.82 6.2a1 1 0 0 0-.32.74v5.56a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1Z"/>`), g.Group(children),
	)
}

func InterfaceHomeThreeHomeHouseMapRoof(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7L7 .5L13.5 7m-11 1.5v5h9v-5"/>`), g.Group(children),
	)
}

func InterfaceHomeTwoDoorEntranceHomeHouseMapRoofRound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 9L7 2.5L13.5 9"/><path d="M2.5 7v6.5h9V7M7 13.5v-4"/></g>`), g.Group(children),
	)
}

func InterfaceIdFaceScanTwoIdentificationAngleSecureHumanIdPersonFaceSecurityBrackets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 10.5v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2m3.5-7v2m6-2v2m-3-1V8H5.5m-1 2a3.63 3.63 0 0 0 5 0"/>`), g.Group(children),
	)
}

func InterfaceIdFingerPrintIdentificationPasswordTouchIdSecureFingerprintFingerSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 8.75v-1.5C13 3.54 10.3.5 7 .5h0c-3.3 0-6 3-6 6.75v6.25m3-2.25v2.25"/><path d="M10 13.5V8c0-2.06-1.35-3.75-3-3.75h0C5.35 4.25 4 5.94 4 8v.75M7 12v1.5m0-6.25V9.5m6 4v-2"/></g>`), g.Group(children),
	)
}

func InterfaceIdFingerPrintScanTouchIdIdentificationHandFingerCircleFingerprint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 13.5V8.25A1.25 1.25 0 0 1 6.75 7h0A1.25 1.25 0 0 1 8 8.25V11h2a2 2 0 0 1 2 2v.5"/><path d="M3.39 8.61a4.75 4.75 0 1 1 6.72 0"/></g>`), g.Group(children),
	)
}

func InterfaceIdIrisScanCheckIdentificationRetinaSecurityApprovedSuccessIrisScanEyeLogin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.5 7.25l1.55 1.14a.48.48 0 0 0 .4.1a.5.5 0 0 0 .34-.24l3.71-6"/><path d="M11.85 5.25a13.92 13.92 0 0 1 1.65 2s-2.91 4.5-6.5 4.5S.5 7.25.5 7.25s2.91-4.5 6.5-4.5"/></g>`), g.Group(children),
	)
}

func InterfaceIdIrisScanIdentificationRetinaSecureSecurityIrisScanEyeBrackets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.5v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2m11-3.5s-2 3-4.5 3s-4.5-3-4.5-3s2-3 4.5-3s4.5 3 4.5 3Z"/><circle cx="7" cy="7" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceIdThumbMarkIdentificationPasswordTouchIdSecureFingerprintFingerSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.39a5 5 0 0 0 5-5V5.61a5 5 0 0 0-1.27-3.33M2 6.72v1.67A5 5 0 0 0 5.06 13M9.5 1.28a5 5 0 0 0-6.83 1.83a4.91 4.91 0 0 0-.57 1.52"/><path d="M6.48 3.51A2.51 2.51 0 0 1 9.5 6v1.61m-.64 2.1A2.5 2.5 0 0 1 4.5 8V6a2.53 2.53 0 0 1 .2-1M7 6.11v1.67"/></g>`), g.Group(children),
	)
}

func InterfaceIdUserIdentificationAngleSecureHumanIdPersonFaceSilhouetteSecurityBrackets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.5v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2"/><circle cx="7" cy="4.5" r="2"/><path d="M10.16 10.5a3.5 3.5 0 0 0-6.32 0"/></g>`), g.Group(children),
	)
}

func InterfaceIdVoiceOneIdentificationSecureIdSoundwaveSoundVoiceSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1.5v11m5.2-9v7M3.1 5v4m10.4-7.5v11m-2.6-9v7M8.3 5v4"/>`), g.Group(children),
	)
}

func InterfaceIdVoiceScanIdentificationSecureIdSoundwaveSoundVoiceBracketsSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5.18v4.14m4-4.14v4.14M5 6.07v2.36m6-3.25v4.14M9 6.07v2.36M.5 4V1.5a1 1 0 0 1 1-1H4M13.5 4V1.5a1 1 0 0 0-1-1H10M.5 10v2.5a1 1 0 0 0 1 1H4m9.5-3.5v2.5a1 1 0 0 1-1 1H10"/>`), g.Group(children),
	)
}

func InterfaceIdVoiceTwoIdentificationSecureIdSoundwaveSoundVoiceBracketsSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4V1.5a1 1 0 0 1 1-1H4M13.5 4V1.5a1 1 0 0 0-1-1H10M.5 10v2.5a1 1 0 0 0 1 1H4m9.5-3.5v2.5a1 1 0 0 1-1 1H10m-9-6h2.5l1.5-4L7 9l2-4l1.5 2.5H13"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderBottomBorderBottomCellFormatFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m-8 13h13m0-8v4m-13-4v4m6.5-4v3M8.5 7h-3m8 0h-2m-9 0h-2M7 .5v2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderCenterBorderCellFormatFormattingHorizontalVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3M7 .5v13M13.5 7H.5m11 6.5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderCornerBorderCellFormatFormattingCornerTopUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m2-9v4m-5 5h-3m-5-1v-11a1 1 0 0 1 1-1h11"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderFrameBorderCellFormatFormattingFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 5.5v3M8.5 7h-3m8 0h-2m-9 0h-2"/><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M7 .5v2m0 9v2"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutBorderFullGridLayoutLayoutsModule(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M7 .5v13M.5 7h13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutBorderHorizontalBorderCellCenterFormatFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3m6.5-3v3M13.5 7H.5m11 6.5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3M7 .5v2m0 9v2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderLeftBorderCellFormatFormattingLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-9-2h4m5 5v3M.5.5v13m6.5-8v3M8.5 7h-3m8 0h-2m0 6.5h1a1 1 0 0 0 1-1v-1m-9 2h4M7 .5v2m0 9v2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderNoneBorderCellFormatFormattingNone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3m6.5-3v3M8.5 7h-3m8 0h-2m-9 0h-2m11 6.5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3M7 .5v2m0 9v2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderRightBorderCellFormatFormattingRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 13.5h-1a1 1 0 0 1-1-1v-1m9 2h-4m-5-5v-3m13 8V.5M7 8.5v-3M5.5 7h3m-8 0h2m0-6.5h-1a1 1 0 0 0-1 1v1m9-2h-4m1.5 13v-2m0-9v-2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderTopBorderCellFormatFormattingTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m2-9v4m-5 5h-3m8-13H.5m8 6.5h-3M7 8.5v-3m0 8v-2m-6.5 0v1a1 1 0 0 0 1 1h1m-2-9v4m13-1.5h-2m-9 0h-2"/>`), g.Group(children),
	)
}

func InterfaceLayoutBorderVerticalBorderCellCenterFormatFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v3m-13-3v3m8-1.5h-3m8 0h-2m-9 0h-2m11 6.5h1a1 1 0 0 0 1-1v-1m-13 0v1a1 1 0 0 0 1 1h1m3 0h3M7 .5v13"/>`), g.Group(children),
	)
}

func InterfaceLayoutEightGridHeaderLayoutLayoutsMasthead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13M7 3.5v10m-6.5-5h13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutElevenColumnLayoutLayoutsLeftSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M4.5.5v13m0-6.5h9"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutFiveColumnHeaderLayoutLayoutsMastheadSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-4 0v10"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutFourColumnHeaderLayoutLayoutsMastheadSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-9 0v10"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutNineColumnLayoutLayoutsLeftSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M4.5.5v13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutOneColumnLayoutLayoutsLeftSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M5.5.5v13m0-6.5h8"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutSevenColumnHeaderLayoutLayoutsMasthead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13M7 3.5v10"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutSixColumnHeaderLayoutLayoutsMasthead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-9 0v10m5-10v10"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutTenColumnLayoutLayoutsRightSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M9.5.5v13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutThreeColumnsColumnsLayoutLayoutsThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M4.5.5v13m5-13v13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutThreeHeaderLayoutLayoutsMastheadTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutTwoColumnHeaderLayoutLayoutsMastheadSidebar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-9 0v10m9-5h-9"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutTwoColumnsColumsLayoutLayoutsTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M7 .5v13"/></g>`), g.Group(children),
	)
}

func InterfaceLayoutTwoRowsRowsLayoutLayoutsTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 7h13"/></g>`), g.Group(children),
	)
}

func InterfaceLightingBrightnessFourBrightAdjustBrightnessAdjustmentSunRaiseControlsDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="3.25"/><path d="M7 .5v1m0 11v1M13.5 7h-1m-11 0h-1m11.1-4.6l-.71.71m-7.78 7.78l-.71.71m9.2 0l-.71-.71M3.11 3.11L2.4 2.4"/></g>`), g.Group(children),
	)
}

func InterfaceLightingBrightnessOneBrightAdjustBrightnessAdjustmentSunRaiseControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="M7 .5v2m-4.6-.1l1.42 1.42M.5 7h2m-.1 4.6l1.42-1.42M7 13.5v-2m4.6.1l-1.42-1.42M13.5 7h-2m.1-4.6l-1.42 1.42"/></g>`), g.Group(children),
	)
}

func InterfaceLightingBrightnessThreeBrightAdjustBrightnessAdjustmentSunRaiseControlsDotSmall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r=".5"/><path d="M7 12.5v-1m3.89-.61l-.71-.71M12.5 7h-1m-.61-3.89l-.71.71M7 1.5v1m-3.89.61l.71.71M1.5 7h1m.61 3.89l.71-.71"/></g>`), g.Group(children),
	)
}

func InterfaceLightingBrightnessTwoBrightAdjustBrightnessAdjustmentSunRaiseControlsHalf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v1m0 11v1M1.5 7h-1m2.61 3.89l-.71.71m.71-8.49L2.4 2.4M7 3.5a3.5 3.5 0 0 0 0 7ZM12.5 7h1m-2.61 3.89l.71.71m-.71-8.49l.71-.71"/>`), g.Group(children),
	)
}

func InterfaceLightingLightBulbLightingLightIncandescentBulbLights(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 5A4.5 4.5 0 1 0 5 9v1.5a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5V9a4.48 4.48 0 0 0 2.5-4ZM5 13.5h4"/>`), g.Group(children),
	)
}

func InterfaceLightingLightBulbOnLightingLightShineIncandescentBulbLights(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8a3.5 3.5 0 1 0-5.06 3.12v1.72a.39.39 0 0 0 .39.38h2.34a.39.39 0 0 0 .39-.38v-1.75A3.5 3.5 0 0 0 10.5 8ZM7 .81v1.5m4-.07L9.94 3.31m3.34 2.03h-1.5M3 2.24l1.06 1.07M.72 5.34h1.5"/>`), g.Group(children),
	)
}

func InterfaceLinkBrokenBreakBrokenHyperlinkLinkRemoveUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.14 7.13L1.27 9a2.65 2.65 0 0 0 0 3.74h0a2.65 2.65 0 0 0 3.74 0l1.11-1.11M9 9.5h1.86a2.64 2.64 0 0 0 2.64-2.64h0a2.64 2.64 0 0 0-2.64-2.64H8.22M7 .5l-.5 2m-6 1l2 1m.5-4l1 2"/>`), g.Group(children),
	)
}

func InterfaceLinkCreateHyperlinkLinkMakeUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 6.49L.79 9.67a1 1 0 0 0 0 1.42l2.12 2.12a1 1 0 0 0 1.42 0L7.51 10M10 7.51l3.18-3.18a1 1 0 0 0 0-1.42L11.09.79a1 1 0 0 0-1.42 0L6.49 4M9 5L5 9"/>`), g.Group(children),
	)
}

func InterfaceLockCircleCircleFrameKeyKeyholeLockLockedSecureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="6" r="1.5"/><path d="M7 7.5v2"/></g>`), g.Group(children),
	)
}

func InterfaceLockCombinationComboLockLockedPadlockSecureSecurityShieldKeyhole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="8" x="2" y="5.5" rx="1"/><path d="M10.5 5.5V4a3.5 3.5 0 0 0-7 0v1.5"/><circle cx="7" cy="9.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceLockShieldCombinationComboLockLockedPadlockSecureSecurityShieldKeyhole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.36 13.4h0a1 1 0 0 1-.72 0h0A9.59 9.59 0 0 1 .5 4.46V1.54a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.92a9.59 9.59 0 0 1-6.14 8.94Z"/><circle cx="7" cy="5.04" r="1.5"/><path d="M7 9.04v-2.5"/></g>`), g.Group(children),
	)
}

func InterfaceLoginArrowEnterFrameLeftLoginPointRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 10v2.5a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1V4m3 3h-8"/><path d="m6.5 5l-2 2l2 2"/></g>`), g.Group(children),
	)
}

func InterfaceLoginCircleArrowEnterLeftLoginPointCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7h-8m2-2l-2 2l2 2"/><path d="M12.48 10.5a6.5 6.5 0 1 1 0-7"/></g>`), g.Group(children),
	)
}

func InterfaceLoginDialPadFingerPasswordDialPadDotFinger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5V8.25A1.25 1.25 0 0 1 8.25 7h0A1.25 1.25 0 0 1 9.5 8.25V11h2a2 2 0 0 1 2 2v.5"/><circle cx="1" cy="1" r=".5"/><circle cx="5" cy="1" r=".5"/><circle cx="9" cy="1" r=".5"/><circle cx="1" cy="4.5" r=".5"/><circle cx="5" cy="4.5" r=".5"/><circle cx="9" cy="4.5" r=".5"/><circle cx="1" cy="8" r=".5"/><path d="M5 8.5a.5.5 0 0 1 0-1Z"/></g>`), g.Group(children),
	)
}

func InterfaceLoginDialPadOneDialPadDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><circle cx="2" cy="7" r="1.5"/><circle cx="7" cy="2" r="1.5"/><circle cx="7" cy="7" r="1.5"/><circle cx="7" cy="12" r="1.5"/><circle cx="12" cy="2" r="1.5"/><circle cx="12" cy="7" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceLoginDialPadTwoDialPadDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><circle cx="2" cy="7" r="1.5"/><circle cx="7" cy="2" r="1.5"/><circle cx="7" cy="7" r="1.5"/><circle cx="7" cy="12" r="1.5"/><circle cx="12" cy="2" r="1.5"/><circle cx="12" cy="7" r="1.5"/><circle cx="2" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceLoginKeyEntryKeyLockLoginPassUnlock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.62 7.38L11.5 1.5l2 2m-4.25.25L11 5.5"/><circle cx="3.5" cy="9.5" r="3"/></g>`), g.Group(children),
	)
}

func InterfaceLoginPasswordLockLoginPadlockPasswordSecureSecurityTextboxType(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y="3.5" rx="1"/><circle cx="3.5" cy="7" r=".5"/><circle cx="6.5" cy="7" r=".5"/><path d="M9.5 8H11"/></g>`), g.Group(children),
	)
}

func InterfaceLogoutArrowExitFrameLeaveLogoutRectangleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 10.5v2a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v2M6.5 7h7m-2-2l2 2l-2 2"/>`), g.Group(children),
	)
}

func InterfaceLogoutCircleArrowEnterRightLogoutPointCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7h7m-2-2l2 2l-2 2"/><path d="M11.7 11.49a6.5 6.5 0 1 1 0-9"/></g>`), g.Group(children),
	)
}

func InterfacePadLockShieldCombinationComboLockLockedPadlockSecureSecurityShieldSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.36 13.43h0a1 1 0 0 1-.72 0h0a9.67 9.67 0 0 1-6.14-9V1.5a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.92a9.67 9.67 0 0 1-6.14 9.01Z"/><rect width="5" height="4" x="4.5" y="5.5" rx="1"/><path d="M8.5 5.5v-1a1.5 1.5 0 1 0-3 0v1"/></g>`), g.Group(children),
	)
}

func InterfacePageControllerButtonLoopOneMultimediaMultiButtonRepeatMediaLoopInfinityControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 10a2.75 2.75 0 0 0 3-3a2.75 2.75 0 0 0-3-3c-2.75 0-4.25 6-7 6a2.75 2.75 0 0 1-3-3a2.75 2.75 0 0 1 3-3c2.75 0 4.25 6 7 6Z"/>`), g.Group(children),
	)
}

func InterfacePageControllerButtonLoopTwoMultimediaMultiButtonRepeatMediaLoopInfinityControls(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 9.43a2.37 2.37 0 0 1-1.5.57a2.75 2.75 0 0 1-3-3a2.75 2.75 0 0 1 3-3c2.75 0 4.25 6 7 6a2.75 2.75 0 0 0 3-3a2.75 2.75 0 0 0-3-3a2.37 2.37 0 0 0-1.45.57"/>`), g.Group(children),
	)
}

func InterfacePageControllerFitScreenFitScreenAdjustDisplayArtboardFrameCorner(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 .5H1a.5.5 0 0 0-.5.5v4m13 0V1a.5.5 0 0 0-.5-.5H9m0 13h4a.5.5 0 0 0 .5-.5V9M.5 9v4a.5.5 0 0 0 .5.5h4"/>`), g.Group(children),
	)
}

func InterfacePageControllerLoadingHalfProgressLoadingLoadHalfWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5a6.5 6.5 0 1 1 6.21-8.41M13.5 7v.5"/><path stroke-dasharray=".889 1.778" d="M13.11 9.23a6.51 6.51 0 0 1-2.79 3.36"/><path d="m9.53 13l-.47.18"/></g>`), g.Group(children),
	)
}

func InterfacePageControllerLoadingOneProgressLoadingLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .88v2.5m-5 0l1.84 1.69M1 8.88l2.42-.9m.97 5.14l1.11-2.24m6.5-7.5l-1.84 1.69M13 8.88l-2.42-.9m-.97 5.14L8.5 10.88"/>`), g.Group(children),
	)
}

func InterfacePageControllerLoadingThreeProgressLoadingDotLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="1.5"/><circle cx="12.25" cy="7" r="1.25"/><circle cx="1.75" cy="7" r="1.25"/></g>`), g.Group(children),
	)
}

func InterfacePageControllerLoadingTwoProgressLoadingLineLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="4" x=".5" y="5" rx=".5"/><rect width="3" height="4" x="8" y="5" rx=".5"/><path d="M13.5 5v4"/></g>`), g.Group(children),
	)
}

func InterfacePageControllerScrollLeftRighMoveScrollHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 2.25l4.39 4.49a.36.36 0 0 1 0 .52L9 11.75m-4-9.5L.61 6.74a.36.36 0 0 0 0 .52L5 11.75"/>`), g.Group(children),
	)
}

func InterfacePageControllerScrollUpDownMoveScrollVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.25 5L6.74.61a.36.36 0 0 1 .52 0L11.75 5m-9.5 4l4.49 4.39a.36.36 0 0 0 .52 0L11.75 9"/>`), g.Group(children),
	)
}

func InterfacePageControllerSettingsPageSettingSquareTriangleCircleLineCombinationVariation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 1h5v5h-5zm12.75 12.5h-5m0-5h5m-5 2.5h5m.25-5H8L10.75.5L13.5 6z"/><circle cx="3" cy="11" r="2.5"/></g>`), g.Group(children),
	)
}

func InterfacePresentationPodiumWorkDeskNotesCompanyPresentationOfficePodiumMicrophone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 8.5H2l-1-5h12l-1 5H9m-2-5v-3m-2 6v7m4-7v7m-5 0h6m-7.5-10s0-3 2-3m7 3s0-3-2-3"/>`), g.Group(children),
	)
}

func InterfaceRemoveCircleDeleteAddCircleSubtractButtonButtonsRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M4 7h6"/></g>`), g.Group(children),
	)
}

func InterfaceRemoveOneButtonDeleteButtonsSubtractHorizontalRemoveLineAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7h13"/>`), g.Group(children),
	)
}

func InterfaceRemoveSquareSubtractButtonsRemoveAddButtonSquareDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 7h6"/><rect width="13" height="13" x=".5" y=".5" rx="3"/></g>`), g.Group(children),
	)
}

func InterfaceRemoveTwoDeleteBoxSubtractButtonsButtonRemoveTextboxTextAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="13" height="3" x=".5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`), g.Group(children),
	)
}

func InterfaceSearchCircleCircleGlassSearchMagnifying(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="6.5" cy="6.5" r="2.5"/><path d="M10 10L8.27 8.27"/></g>`), g.Group(children),
	)
}

func InterfaceSearchGlassSearchMagnifying(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.92" cy="5.92" r="5.42"/><path d="M13.5 13.5L9.75 9.75"/></g>`), g.Group(children),
	)
}

func InterfaceSearchSquareGlassSearchSquareMagnifying(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.5" cy="6.5" r="2.5"/><path d="M10 10L8.27 8.27"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceSecurityFireWallCodeFirewallProgrammingSecureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 7.5h13v6H.5z"/><path d="M7 .5S8.5 3 6.5 4C5.5 4.5 4 3 4 3a3.18 3.18 0 0 0 3 4.5A3.34 3.34 0 0 0 10.5 4A3.34 3.34 0 0 0 7 .5Zm-6.5 10h13m-8 0v3m3-6v3"/></g>`), g.Group(children),
	)
}

func InterfaceSecurityShieldFourShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.36 13.43h0a1 1 0 0 1-.72 0h0A9.57 9.57 0 0 1 .5 4.49V1.88a.51.51 0 0 1 .36-.49a21.57 21.57 0 0 1 12.28 0a.51.51 0 0 1 .36.49v2.61a9.57 9.57 0 0 1-6.14 8.94Z"/>`), g.Group(children),
	)
}

func InterfaceSecurityShieldOneShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.36 13.43h0a1 1 0 0 1-.72 0h0A9.57 9.57 0 0 1 .5 4.49v-3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v3a9.57 9.57 0 0 1-6.14 8.94Z"/>`), g.Group(children),
	)
}

func InterfaceSecurityShieldPersonshieldSecureSecurityPerson(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M6 6.61A4.49 4.49 0 0 0 .5 11v1.5H6m4.67.97h0a.5.5 0 0 1-.34 0h0A4.48 4.48 0 0 1 7.5 9.31V8a.47.47 0 0 1 .5-.5h5a.47.47 0 0 1 .5.5v1.31a4.48 4.48 0 0 1-2.83 4.16Z"/></g>`), g.Group(children),
	)
}

func InterfaceSecurityShieldProfileshieldSecureSecurityProfilePerson(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.36 13.43h0a1 1 0 0 1-.72 0h0a9.67 9.67 0 0 1-6.14-9V1.5a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.92a9.67 9.67 0 0 1-6.14 9.01Z"/><circle cx="7" cy="5.5" r="2.5"/><path d="M3.25 11.19a5 5 0 0 1 7.5 0"/></g>`), g.Group(children),
	)
}

func InterfaceSecurityShieldThreeShieldPayProductSecureMoneyCashCurrencySecurityBusiness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 6.94a5.64 5.64 0 0 1-3.13 5L7 13.5L4.13 12A5.64 5.64 0 0 1 1 6.94V1.5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/><path d="M5.5 8A1.5 1.5 0 0 0 7 9.5A1.4 1.4 0 0 0 8.5 8c0-1.5-3-1.5-3-3A1.4 1.4 0 0 1 7 3.5A1.5 1.5 0 0 1 8.5 5M7 3.5v-1m0 8v-1"/></g>`), g.Group(children),
	)
}

func InterfaceSecurityShieldTwoShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.92 13.21l-.59.23a.94.94 0 0 1-.66 0l-.59-.23A8 8 0 0 1 1 5.78V3A6.36 6.36 0 0 0 7 .5c1.25 1.82 3.32 2.68 6 2.5v2.78a8 8 0 0 1-5.08 7.43Z"/>`), g.Group(children),
	)
}

func InterfaceSettingCogWorkLoadingCogGearSettingsMachine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.23 2.25l.43-1.11A1 1 0 0 1 6.59.5h.82a1 1 0 0 1 .93.64l.43 1.11l1.46.84l1.18-.18a1 1 0 0 1 1 .49l.4.7a1 1 0 0 1-.08 1.13l-.73.93v1.68l.75.93a1 1 0 0 1 .08 1.13l-.4.7a1 1 0 0 1-1 .49l-1.18-.18l-1.46.84l-.43 1.11a1 1 0 0 1-.93.64h-.84a1 1 0 0 1-.93-.64l-.43-1.11l-1.46-.84l-1.18.18a1 1 0 0 1-1-.49l-.4-.7a1 1 0 0 1 .08-1.13L2 7.84V6.16l-.75-.93a1 1 0 0 1-.08-1.13l.4-.7a1 1 0 0 1 1-.49l1.18.18ZM5 7a2 2 0 1 0 2-2a2 2 0 0 0-2 2Z"/>`), g.Group(children),
	)
}

func InterfaceSettingGaugeDashboardOneBarSpeedTestLoadingDashboardInternetGaugeProgress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m10 4.5l-3 5m-6 0h12"/></g>`), g.Group(children),
	)
}

func InterfaceSettingGaugeDashboardTwoBarSpeedTestLoadingDashboardInternetGaugeProgress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 6.5l-3 5m6 0a6.5 6.5 0 1 0-12 0Z"/>`), g.Group(children),
	)
}

func InterfaceSettingHammerConstructionHammerMalletToolToolsSettingEditAdjust(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5.66" height="9.66" x="5.67" y=".67" rx="1" transform="rotate(-45 8.498 5.5)"/><path d="M.91 11.09a1.41 1.41 0 0 0 2 2L7.5 8.5l-2-2Z"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuHorizontalCircleNavigationDotsThreeCircleButtonHorizontalMenu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r=".5"/><circle cx="4" cy="7" r=".5"/><circle cx="10" cy="7" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuHorizontalNavigationDotsThreeCircleButtonHorizontalMenu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="7" r="1.5"/><circle cx="7" cy="7" r="1.5"/><circle cx="2" cy="7" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuHorizontalSquareNavigationDotsThreeSquareButtonHorizontalMenu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="7" cy="7" r=".5"/><circle cx="4" cy="7" r=".5"/><circle cx="10" cy="7" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuOneAlternateButtonParallelHorizontalLinesMenuNavigationTwoThick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="4" x="1.5" y="9.61" rx="1"/><rect width="11" height="4" x="1.5" y=".61" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuOneButtonParallelHorizontalLinesMenuNavigationThreeHamburger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2H.5m13 5H.5m13 5H.5"/>`), g.Group(children),
	)
}

func InterfaceSettingMenuParallelHamburgerCircleNavigationParallelHamburgerButtonmenuCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M4.5 4.5h5M4.5 7h5m-5 2.5h5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuParallelHamburgerSquareNavigationParallelHamburgerButtonmenuSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M4.5 4.5h5M4.5 7h5m-5 2.5h5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuTwoAlternateButtonParallelHorizontalLinesMenuNavigationStaggeredThick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="4" x=".5" y="1.5" rx="1"/><rect width="9" height="4" x="4.5" y="8.5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceSettingMenuTwoButtonParallelHorizontalLinesMenuNavigationStaggeredThreeHamburger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2H6m5 5H3.5m5 5h-8"/>`), g.Group(children),
	)
}

func InterfaceSettingMenuVerticalNavigationVerticalThreeCircleButtonMenuDots(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2" r="1.5"/><circle cx="7" cy="7" r="1.5"/><circle cx="7" cy="12" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingPieChartCogSettingGraphCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m5.23 2.25l.43-1.11A1 1 0 0 1 6.59.5h.82a1 1 0 0 1 .93.64l.43 1.11l1.46.84l1.18-.18a1 1 0 0 1 1 .49l.4.7a1 1 0 0 1-.08 1.13l-.73.93v1.68l.75.93a1 1 0 0 1 .08 1.13l-.4.7a1 1 0 0 1-1 .49l-1.18-.18l-1.46.84l-.43 1.11a1 1 0 0 1-.93.64h-.84a1 1 0 0 1-.93-.64l-.43-1.11l-1.46-.84l-1.18.18a1 1 0 0 1-1-.49l-.4-.7a1 1 0 0 1 .08-1.13L2 7.84V6.16l-.75-.93a1 1 0 0 1-.08-1.13l.4-.7a1 1 0 0 1 1-.49l1.18.18Z"/><circle cx="7" cy="7" r="3"/><path d="M7 4v3h3M7 7L4.88 9.12"/></g>`), g.Group(children),
	)
}

func InterfaceSettingScrewdriverScrewdriverScrewToolSettingsHand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.85 7.037L6.973 9.16a.49.49 0 0 1 0 .693l-3.125 3.125a2 2 0 0 1-2.829 0h0a2 2 0 0 1 0-2.828l3.126-3.126a.49.49 0 0 1 .707.014ZM5.91 8.09l4.48-4.48m.35.34l-.69-.69V1.88L12.12.5l1.38 1.38l-1.38 2.07h-1.38z"/>`), g.Group(children),
	)
}

func InterfaceSettingSlideHorizontalAdjustmentAdjustControlsFaderHorizontalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.5" cy="7" r="2"/><path d="M1.5 7h-1m9 0h4"/></g>`), g.Group(children),
	)
}

func InterfaceSettingSliderHorizontalAdjustmentAdjustControlsFaderHorizontalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><path d="M3.5 2h10"/><circle cx="7" cy="7" r="1.5"/><path d="M.5 7h5m3 0h5"/><circle cx="12" cy="12" r="1.5"/><path d="M10.5 12H.5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingSliderHorizontalSquareAdjustControlsFaderHorizontalSettingsSliderSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M3 4.5h5"/><circle cx="9.5" cy="4.5" r="1.5"/><path d="M11 9.5H8m-3 0H3"/><circle cx="6.5" cy="9.5" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="4.5" r="1.5"/><path d="M2 6v7.5m0-13V3"/><circle cx="12" cy="4.5" r="1.5"/><path d="M12 3V.5m0 13V6"/><circle cx="7" cy="7" r="1.5"/><path d="M7 .5v5m0 3v5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSliderOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2"/><path d="M7 1.5v-1m0 9v4"/></g>`), g.Group(children),
	)
}

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSliderSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(-90 7 7)"/><path d="M4.5 11V6"/><circle cx="4.5" cy="4.5" r="1.5"/><path d="M9.5 3v3m0 3v2"/><circle cx="9.5" cy="7.5" r="1.5"/></g>`), g.Group(children),
	)
}

func InterfaceSettingToolBoxBoxBriefcaseToolSettings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="4.5" rx="1"/><path d="M.5 8.5h13M7 7.5v2m3-5a3 3 0 0 0-3-3h0a3 3 0 0 0-3 3"/></g>`), g.Group(children),
	)
}

func InterfaceSettingWrenchCrescentToolConstructionToolsWrenchSettingEditAdjust(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.43 3.59a.76.76 0 0 0-.35-.51l-2 2a1 1 0 0 1-1.44 0l-.76-.68a1 1 0 0 1 0-1.4l2-2a.76.76 0 0 0-.48-.43A3.8 3.8 0 0 0 6.26 6L.8 11.41a1 1 0 0 0 0 1.43l.36.36a1 1 0 0 0 1.43 0l5.46-5.45a3.81 3.81 0 0 0 5.38-4.16Z"/>`), g.Group(children),
	)
}

func InterfaceSettingZoomAreaZoomMagnifierSquareArea(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v2.75M.5 5.5v3m0 3v1a1 1 0 0 0 1 1h1m3 0h2.75"/><circle cx="8" cy="8" r="3.5"/><path d="M10.47 10.47L13 13"/></g>`), g.Group(children),
	)
}

func InterfaceShareHandLockGiveHandLockPadlockSecureSecurityTransfer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6" height="4.5" x="7.5" y="3.5" rx=".5"/><path d="M8.5 3.5v-1a2 2 0 0 1 4 0v1m-12 5h3A2.5 2.5 0 0 1 6 11h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5h-9"/></g>`), g.Group(children),
	)
}

func InterfaceShareMegaPhoneOneBullhornLoudMegaphoneShareSpeakerTransmit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.84 12.07L1.17 8a1 1 0 0 1-.67-.91V6a1 1 0 0 1 .67-.94L12.84 1a.5.5 0 0 1 .66.47V11.6a.5.5 0 0 1-.66.47Zm-4.36-1.5A2.75 2.75 0 0 1 3 10.25V8.66"/>`), g.Group(children),
	)
}

func InterfaceShareSatelliteBroadcastSatelliteShareTransmit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 9A6 6 0 0 1 5 .5ZM9.26 4.74L12 2"/><path d="M3.96 7.57L.5 13.5H7L5.92 9.73"/></g>`), g.Group(children),
	)
}

func InterfaceShareShareTransmit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.75" cy="7" r="2.25"/><circle cx="11.25" cy="11.25" r="2.25"/><circle cx="11.25" cy="2.75" r="2.25"/><path d="m4.76 6l4.48-2.25M4.76 8l4.48 2.25"/></g>`), g.Group(children),
	)
}

func InterfaceShareUserHumanPersonShareSignalTransmitUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.75" r="2.25"/><path d="M11 13.5a4.5 4.5 0 0 0-8 0"/><path d="M12 10.56a6.25 6.25 0 1 0-9.92 0"/></g>`), g.Group(children),
	)
}

func InterfaceSignalGraphCircleCircleStatsGraphLineBeatHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7.01h4L6 4.5L7.5 10l2-2.99h4M12.77 4A6.51 6.51 0 0 0 1.23 4m0 6a6.51 6.51 0 0 0 11.54 0"/>`), g.Group(children),
	)
}

func InterfaceSignalGraphHeartLineBeatSquareGraphStats(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7.08h2.19a.52.52 0 0 0 .45-.27l1.8-3.6a.49.49 0 0 1 .49-.27a.48.48 0 0 1 .43.35l2.23 7.42a.5.5 0 0 0 .46.36a.5.5 0 0 0 .45-.32l1.37-3.35a.51.51 0 0 1 .47-.32h2.66"/>`), g.Group(children),
	)
}

func InterfaceSignalSquareHeartLineStatsBeatSquareGraph(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="M2.5 7.02h2L6 4.51l1.5 5.5l2-2.99h2"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingBoldTextBoldFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 3.5a3 3 0 0 0-3-3H4a1 1 0 0 0-1 1v5h5a3 3 0 0 0 3-3Zm1 6.5a3.5 3.5 0 0 1-3.5 3.5H4a1 1 0 0 1-1-1v-6h5.5A3.5 3.5 0 0 1 12 10Z"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingCenterAlignTextAlignmentAlignParagraphCenteredFormattingCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1h13M2 4h10M3.5 7h7m-10 6h13M2 10h10"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingEraserTextRemoveFormatFormattingEraserDelete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 13.5h10m-1.18-7.29a.94.94 0 0 0 0-1.35L8.25.78a1 1 0 0 0-1.36 0L.78 6.89a1 1 0 0 0 0 1.36l2.44 2.43a.92.92 0 0 0 .67.29h3.28a.92.92 0 0 0 .68-.29Z"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingFilterOneFunnelFilterRoundOil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5H.5a6.51 6.51 0 0 0 5 6.33v6.67l3-2V6.83a6.51 6.51 0 0 0 5-6.33Z"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingFontSizeSizeTextFormattingFontFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 12.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33M7.91 9.21h4.18M.5 12.5L5.07 1.84a.57.57 0 0 1 1 0L7 3.94M2.55 7.72H5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingImageBottomInsertParagraphBottomImageTextAlignFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 6.5h-9m13-3H.5m13-3H.5"/><rect width="4" height="13" x="5" y="5" rx=".5" transform="rotate(-90 7 11.5)"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingImageLeftAlignmentWrapFormattingParagraphImageLeftText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 1H9m4.5 3H9m4.5 3H9m4.5 6H.5m13-3H.5"/><rect width="6" height="6" x=".5" y="1" rx=".5"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingImageRightParagraphImageTextAlignmentWrapRightFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 1H5M.5 4H5M.5 7H5M.5 13h13m-13-3h13"/><rect width="6" height="6" x="7.5" y="1" rx=".5" transform="rotate(-180 10.5 4)"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingImageTopInsertParagraphImageTextAlignTopFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 13.5h-9m13-3H.5m13-3H.5"/><rect width="4" height="13" x="5" y="-4" rx=".5" transform="rotate(-90 7 2.5)"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingIndentDecreaseTextAlignmentIndentParagraphAlignFormattingDecrease(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 4H6m7.5 4H6m7.5 4H.5m2-8L1 7l1.5 2"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingIndentIncreaseAlignmentAlignIndentParagraphIncreaseFormattingText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 4H6m7.5 4H6m7.5 4H.5M2 5l2 2l-2 2"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingIndentLeftAlignmentAlignIndentParagraphFormattingLeftText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1.5h-6m4 4h-4m6 4h-6m4 4h-4M.5 5l2 2l-2 2M5 13.5V.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingIndentRightAlignmentAlignIndentParagraphFormattingRightText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1.5h6m-4 4h4m-6 4h6m-4 4h4m7-8.5l-2 2l2 2M9 13.5V.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingItalicOffTextOffFormattingItalicFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.24 6.24L4 13.5m-2.5 0h5M5 .5h7.5a1 1 0 0 1 1 1V3M.5.5l13 13M8 .5L7.23 3"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingItalicTextFormattingItalicFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 13.5l5-13m-4 0H13m-12 13h7.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingJustifiedAlignJustifiedAlignAlignmentParagraphFormattingText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 3H.5m13 3H.5m13 6H.5m13-3H.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingLeftAlignParagraphTextAlignmentAlignLeftFormattingRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1h13M.5 4h10M.5 7h7m-7 6h13m-13-3h10"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingListBulletsPointsBulletUnorderedListListsBullets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="1" cy="2.5" r=".5"/><path d="M4.5 2.5h9"/><circle cx="1" cy="7" r=".5"/><path d="M4.5 7h9"/><circle cx="1" cy="11.5" r=".5"/><path d="M4.5 11.5h9"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingParagraphAlignmentParagraphFormattingText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5h-9a4 4 0 0 0 0 8h2m3-8v13m-3-13v13"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingParagraphArticleAlignmentFormattingNormalParagraphText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H5m8.5 3H.5m13 3H.5m9 6h-9m13-3H.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingParagraphBulletsPointsBulletAlignParagraphFormattingBulletsText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 1H8m5.5 3H8m5.5 3H8m5.5 6H8m5.5-3H8"/><rect width="5" height="5" x=".5" y="1" rx=".5" transform="rotate(-90 3 3.5)"/><rect width="5" height="5" x=".5" y="8" rx=".5" transform="rotate(-90 3 10.5)"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingParagraphColumnTextAlignmentFormattingParagraphColumnsColumnTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 1H2m3.5 3h-5m5 3h-5m5 6h-5m5-3h-5m13-9h-5m5 3h-5m5 3h-5m3.5 6H8.5m5-3h-5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingQuotationFourQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4L1 7m2-3l.5 3m7.75 0l-.5 3m2.75-3l-.5 3"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingQuotationOneQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.07 7a3.82 3.82 0 0 1 0-4m3.01 4a3.82 3.82 0 0 1 0-4m8.85 8a3.82 3.82 0 0 0 0-4m-3.01 4a3.82 3.82 0 0 0 0-4"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingQuotationThreeQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.08 7.25a3.82 3.82 0 0 1 0-4"/><circle cx="4.45" cy="6.92" r=".5"/><path d="M1.07 7.25a3.82 3.82 0 0 1 0-4"/><circle cx="1.44" cy="6.92" r=".5"/><path d="M9.92 6.75a3.82 3.82 0 0 1 0 4"/><circle cx="9.55" cy="7.08" r=".5"/><path d="M12.93 6.75a3.82 3.82 0 0 1 0 4"/><circle cx="12.56" cy="7.08" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingQuotationTwoQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.18 5.44a1 1 0 0 1 .42.81a.9.9 0 0 1-1.1 1C.54 7 .09 5.57 1 3.25m4.7 2.19a1 1 0 0 1 .43.81a.9.9 0 0 1-1.1 1c-1-.19-1.41-1.66-.53-4m7.32 5.31a1 1 0 0 1-.42-.81a.9.9 0 0 1 1.1-1c1 .19 1.41 1.66.52 4M8.3 8.56a1 1 0 0 1-.43-.81a.9.9 0 0 1 1.1-1c1 .19 1.41 1.66.53 4"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingRightAlignRagParagraphTextAlignmentAlignRightFormattingLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 3h-10m10 3h-7m7 6H.5m13-3h-10"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingSigmaFormulaTextFormatSigmaFormattingSum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5H1a.5.5 0 0 1-.46-.31a.47.47 0 0 1 .11-.54l5.29-5.3a.5.5 0 0 0 0-.7L.65 1.35A.47.47 0 0 1 .54.81A.5.5 0 0 1 1 .5h12.5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingStrikeThroughTextStrikeThroughFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.08 8.67a2.73 2.73 0 0 1-.42 3.95a5.06 5.06 0 0 1-5.37.21l-.9-.51m6.61-11C8 .33 4.23-.26 3.38 2.54a2.47 2.47 0 0 0 .55 2.38a11.8 11.8 0 0 0 2.6 1.54M.5 6.5h13"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingSubScriptTextFormattingSubscriptFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1 .5l7 8m0-8l-7 8m9.5 2.5v-.25a1.25 1.25 0 0 1 1.25-1.25h0A1.25 1.25 0 0 1 13 10.75a1.58 1.58 0 0 1-.59 1.25l-1.91 1.5H13"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingSuperScriptTextFormattingSuperscriptFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1 5.5l7 8m0-8l-7 8M10.5 2v-.25A1.25 1.25 0 0 1 11.75.5h0A1.25 1.25 0 0 1 13 1.75A1.58 1.58 0 0 1 12.41 3L10.5 4.5H13"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextBarTextBarFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 13.5H9a2 2 0 0 1-2-2a2 2 0 0 1-2 2H3.5m7-13H9a2 2 0 0 0-2 2a2 2 0 0 0-2-2H3.5m3.5 2v9M4.5 8h5"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextSquareTextOptionsFormattingFormatSquareColorBorderFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 10.78l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33M4.69 8h4.62"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextStrikeThroughTextStrikeThroughFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5V5m-2.5 8.5h5M1.5 3V1.5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1V3m-11 4.5h11M7 7.5v6"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextStyleTextStyleFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v13m-2.5 0h5M1.5 3V1.5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1V3"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextTextOptionsFormattingFormatColor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5h13M2.14 11L6.5.83a.54.54 0 0 1 1 0L11.86 11M4.07 6.5h5.86"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextToSpeechOneSpeechAutomatedTranslateLanguageVoiceTechnology(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 13.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33m-5.59-3.29h4.18M.5 4A3.5 3.5 0 0 1 4 .5"/><path d="M.5 4A3.5 3.5 0 0 1 4 .5M3 4a1 1 0 0 1 1-1"/><path d="M3 4a1 1 0 0 1 1-1m6-2.5A3.5 3.5 0 0 1 13.5 4"/><path d="M10 .5A3.5 3.5 0 0 1 13.5 4M10 3a1 1 0 0 1 1 1"/><path d="M10 3a1 1 0 0 1 1 1"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingTextToSpeechTwoSpeechAutomatedTranslateLanguageVoiceTechnology(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 7.5h8m-4-2v2m-2.5 0c0 2.32 2.16 4.28 5 6"/><path d="M9.5 7.5c.05 2.32-2.16 4.28-5 6M.5 4A3.5 3.5 0 0 1 4 .5"/><path d="M.5 4A3.5 3.5 0 0 1 4 .5M3 4a1 1 0 0 1 1-1"/><path d="M3 4a1 1 0 0 1 1-1m6-2.5A3.5 3.5 0 0 1 13.5 4"/><path d="M10 .5A3.5 3.5 0 0 1 13.5 4M10 3a1 1 0 0 1 1 1"/><path d="M10 3a1 1 0 0 1 1 1"/></g>`), g.Group(children),
	)
}

func InterfaceTextFormattingTranslateOptionsTextTranslate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 13.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33m-5.59-3.29h4.18M.5 2.5h9M5 .5v2m3 0l-6 7m0-7l3.94 4.6"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingTypeCursorCursorTextFormattingTypeFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5H13a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H7.5m-5 0H1a.5.5 0 0 1-.5-.5v-3A.5.5 0 0 1 1 5h1.5M4 2.5h2m-2 9h2m-1-9v9"/>`), g.Group(children),
	)
}

func InterfaceTextFormattingUnderlineTextUnderlineFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v6.75a4.25 4.25 0 0 0 4.25 4.25h0A4.25 4.25 0 0 0 11 7.25V.5m-9.5 13h11"/>`), g.Group(children),
	)
}

func InterfaceTimeAlarmNotificationAlertBellWakeClockAlarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="8" r="5.5"/><path d="M5.5.5h3M7 .5v2M5.5 6L7 8h2.5M12 2l1 1M2 2L1 3"/></g>`), g.Group(children),
	)
}

func InterfaceTimeClockCircleClockLoadingMeasureTimeCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 4.5V7l2.54 2.96"/></g>`), g.Group(children),
	)
}

func InterfaceTimeClockHandHandClockTimeGiveHumanPerson(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.28 5.25a4.5 4.5 0 1 1 7.93 3.38"/><path d="M8.75 3.75v2h1.5m-9.5 2h4a2.5 2.5 0 0 1 2.5 2.5h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5h-10"/></g>`), g.Group(children),
	)
}

func InterfaceTimeClockNineToFiveWorkHourTimeClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 12v1.5m-3.54-2.96L2.4 11.6M7 .5A6.5 6.5 0 0 0 .5 7H7l4.6 4.6A6.51 6.51 0 0 0 7 .5Z"/>`), g.Group(children),
	)
}

func InterfaceTimeClockSquareClockLoadingFrameMeasureTimeCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="4"/><path d="M7 5.5V7h1.5"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceTimeHourGlassHourglassLoadingMeasureClockTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 3.5a3.5 3.5 0 0 1-7 0v-3h7Z"/><path d="M10.5 10.5a3.5 3.5 0 0 0-7 0v3h7Zm-9-10h11m-11 13h11"/></g>`), g.Group(children),
	)
}

func InterfaceTimeMidnightWholeMidnightHourClockTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6M12 7h1.5m-2.96 3.54l1.06 1.06M7 12v1.5m3.54-10.04L11.6 2.4M7 2V.5"/>`), g.Group(children),
	)
}

func InterfaceTimeNineThreeQuartersFourtyFiveClockMinutesNineTimeHour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7A6.5 6.5 0 1 0 7 .5V7Zm2.96-3.54L2.4 2.4"/>`), g.Group(children),
	)
}

func InterfaceTimeResetTimeClockResetStopwatchCircleMeasureLoading(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 .5A6.5 6.5 0 1 1 .5 7a7.23 7.23 0 0 1 2-5"/><path d="m.5 2.5l2-.5l.5 2m4-.5v4l2.6 1.3"/></g>`), g.Group(children),
	)
}

func InterfaceTimeRewindBackReturnClockTimerCountdown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 7A6.5 6.5 0 1 0 7 .5a7.23 7.23 0 0 0-5 2"/><path d="m2.5.5l-.5 2L4 3m3 .5v4L4.4 8.8"/></g>`), g.Group(children),
	)
}

func InterfaceTimeSixHourClockTimeMinutesHalfThirtySix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5a6.5 6.5 0 0 0 0-13ZM2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6"/>`), g.Group(children),
	)
}

func InterfaceTimeSleepNapSleepRestBreakClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7A6.5 6.5 0 1 1 7 .5"/><path d="M7 4.5V7L4.46 9.96M10 .5h3.5l-3.5 4h3.5"/></g>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchAlternateTimerCountdownClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5A6.5 6.5 0 1 1 13.5 7a7.23 7.23 0 0 1-2 5"/><path d="m13.5 11.5l-2 .5l-.5-2M9 9L7 6.5H4"/></g>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchHalfQuarterSecondsThirtyHourHalfMeasureTimeStopwatch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.75 2.5a5.5 5.5 0 0 1 0 11Zm-1.5-2h3m-1.5 0v2m5-.5l1 1M2.86 4.11l1.06 1.06M1.25 8h1.5m.11 3.89l1.06-1.06"/>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchQuarterTimeFifteenQuarterStopwatchMinuteMeasureSeconds(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.75 8V2.5a5.5 5.5 0 0 1 5.5 5.5ZM5.25.5h3m-1.5 0v2m5-.5l1 1M2.86 4.11l1.06 1.06M1.25 8h1.5m.11 3.89l1.06-1.06M7 12v1.5m3.54-2.96l1.06 1.06"/>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchSixtySecondMinuteTimeStopwatchMeasureHourWholeClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5.5h3M7 .5V4m-3.89.11l1.06 1.06M1.5 8H3m.11 3.89l1.06-1.06M7 13.5V12m3.89-.11l-1.06-1.06M12.5 8H11m-.11-3.89L9.83 5.17"/>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchThirdQuarterQuartersTimeMeasureStopwatchThreeSecondsMinutes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.75 2.5A5.5 5.5 0 1 1 1.25 8h5.5Zm-1.5-2h3m-1.5 2v-2m5 1.5l1 1M2.86 4.11l1.06 1.06"/>`), g.Group(children),
	)
}

func InterfaceTimeStopWatchTimerCountdownClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 6.92a7.62 7.62 0 0 1-.74 3.21M7 13.42a6.5 6.5 0 0 1-6.5-6.5m8.5 2l-2-2.5H4"/><circle cx="12" cy="11.42" r="1.5"/><path d="M.92 4.52A6.31 6.31 0 0 1 1.79 3m1.93-1.72a6.3 6.3 0 0 1 1.81-.7m7.55 3.94A6.31 6.31 0 0 0 12.21 3m-1.93-1.72a6.3 6.3 0 0 0-1.81-.7"/></g>`), g.Group(children),
	)
}

func InterfaceTimeThreeClockFifteenThreeQuarterMinutesTimeHour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7A6.5 6.5 0 0 0 7 .5V7ZM2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6M7 12v1.5m3.54-2.96l1.06 1.06"/>`), g.Group(children),
	)
}

func InterfaceTimeTimerTimeTockStopwatchMeasureClockTick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="8" r="5.5"/><path d="M.5 2.5A8.69 8.69 0 0 1 3 .5m10.5 2a8.69 8.69 0 0 0-2.5-2M7 5v3h2.5"/></g>`), g.Group(children),
	)
}

func InterfaceUnlinkBreakBrokenHyperlinkLinkRemoveUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.54 9.7l2-2l2.12-2.12a3 3 0 0 0 0-4.24h0a3 3 0 0 0-4.24 0L7 2.76L5.46 4.3M3.5 6.26L1.38 8.38a3 3 0 0 0 0 4.24h0a3 3 0 0 0 4.24 0l1-1M4.17 2.05l5.66 9.9"/>`), g.Group(children),
	)
}

func InterfaceUnlockCombinationComboKeyKeyholeLockSecureSecuritySquareUnlockUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="8" x="2" y="5.5" rx="1"/><path d="M9.47 1.53A3.44 3.44 0 0 0 7 .5A3.5 3.5 0 0 0 3.5 4v1.5"/><circle cx="7" cy="9.5" r=".5"/></g>`), g.Group(children),
	)
}

func InterfaceUploadBoxOneArrowBoxDownloadInternetNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 13.5h-2a1 1 0 0 1-1-1v-8h13v8a1 1 0 0 1-1 1h-2"/><path d="M4.5 10L7 7.5L9.5 10M7 7.5v6M11.29 1a1 1 0 0 0-.84-.5h-6.9a1 1 0 0 0-.84.5L.5 4.5h13ZM7 .5v4"/></g>`), g.Group(children),
	)
}

func InterfaceUploadButtonOneArrowButtonDownloadInternetNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 5h1a.5.5 0 0 1 .5.5V13a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 2 13V5.5a.5.5 0 0 1 .5-.5h1M7 7.5v-7m-2 2l2-2l2 2"/>`), g.Group(children),
	)
}

func InterfaceUploadButtonTwoArrowBottomDownloadInternetNetworkErverUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 10.5v1a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M4 4L7 .5L10 4M7 .5v9"/>`), g.Group(children),
	)
}

func InterfaceUploadCircleArrowCircleDownloadInternetNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4 7l3-3.5L10 7M7 3.5v7"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceUploadDesktopActionActionsComputerDesktopDeviceDisplayMonitorScreenUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 3h1a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-7A.5.5 0 0 1 1 3h1m5 8v2.5m-2 0h4m-2-7v-6m-2 2l2-2l2 2"/>`), g.Group(children),
	)
}

func InterfaceUploadLaptopArrowComputerDownloadInternetLaptopNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.91 9.5l-1.3 2.55a1 1 0 0 0 0 1a1 1 0 0 0 .87.47h11a1 1 0 0 0 .87-.47a1 1 0 0 0 0-1L12.09 9.5ZM5 2.5l2-2l2 2m-2-2v6"/><path d="M3 4.5a1 1 0 0 0-1 1v4h10v-4a1 1 0 0 0-1-1"/></g>`), g.Group(children),
	)
}

func InterfaceUploadSquareArrowDownloadInternetNetworkServerSquareUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m4 7l3-3.5L10 7M7 3.5v7"/></g>`), g.Group(children),
	)
}

func InterfaceUploadWebsiteActionActionsComputerWebsiteDeviceDisplayUploadMonitorScreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 13.5h-2a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-2m-10-10h13M7 13.5V7"/><path d="M4.5 9.5L7 7l2.5 2.5"/></g>`), g.Group(children),
	)
}

func InterfaceUserAddActionsAddCloseGeometricHumanPersonPlusSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M4.5 12.5h-4V11A4.51 4.51 0 0 1 7 7m3.5.5v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func InterfaceUserBlockActionsBlockCloseDeniedDenyGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><circle cx="10.25" cy="10.25" r="3.25"/><path d="m7.95 12.55l4.6-4.6M6 6.61A4.49 4.49 0 0 0 .5 11v1.5h4"/></g>`), g.Group(children),
	)
}

func InterfaceUserCheckActionsCloseCheckmarkCheckGeometricHumanPersonSingleSuccessUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 8l-4.12 5.5l-2.75-2.06"/><circle cx="5" cy="2.75" r="2.25"/><path d="M3 12.5H.5V11a4.5 4.5 0 0 1 7.68-3.18"/></g>`), g.Group(children),
	)
}

func InterfaceUserCircleCircleGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2.5"/><path d="M2.73 11.9a5 5 0 0 1 8.54 0"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceUserDeleteActionsCloseDeleteDenyFailGeometricHumanPersonRemoveSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M5 12.5H.5V11a4.5 4.5 0 0 1 6.73-3.91m6.27 2.17L9.26 13.5m0-4.24l4.24 4.24"/></g>`), g.Group(children),
	)
}

func InterfaceUserEditActionsCloseEditGeometricHumanPencilPersonSingleUpUserWrite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M3.5 12.5h-3V11A4.51 4.51 0 0 1 7 7m6.5 1.5l-4.71 4.71l-2.13.29l.3-2.13l4.7-4.71L13.5 8.5z"/></g>`), g.Group(children),
	)
}

func InterfaceUserFullBodyGeometricHumanPersonSingleUpUserFullBody(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M10.5 8a3.5 3.5 0 0 0-7 0v1.5H5l.5 4h3l.5-4h1.5Z"/></g>`), g.Group(children),
	)
}

func InterfaceUserHomeHomeGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="M2.73 13.4a5 5 0 0 1 8.54 0"/><path d="M13.5 6.94a1 1 0 0 0-.32-.74L7 .5L.82 6.2a1 1 0 0 0-.32.74v5.56a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1Z"/></g>`), g.Group(children),
	)
}

func InterfaceUserHumanGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 13.5V11h2a1 1 0 0 0 1-1V6a5.5 5.5 0 1 0-8 4.9v2.6"/>`), g.Group(children),
	)
}

func InterfaceUserLockActionsLockGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M6 6.61A4.49 4.49 0 0 0 .5 11v1.5h4"/><rect width="6" height="4.5" x="7.5" y="9" rx=".5"/><path d="M8.5 9V8a2 2 0 0 1 4 0v1"/></g>`), g.Group(children),
	)
}

func InterfaceUserMultipleCloseGeometricHumanMultiplePersonUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="3.75" r="2.25"/><path d="M9.5 13.5h-9v-1a4.5 4.5 0 0 1 9 0ZM9 1.5A2.25 2.25 0 0 1 9 6m1.6 2.19a4.5 4.5 0 0 1 2.9 4.2v1.11H12"/></g>`), g.Group(children),
	)
}

func InterfaceUserProfileFocusCloseGeometricHumanPersonProfileFocusUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2.5"/><path d="M10.31 10.75a5 5 0 0 0-6.62 0m9.81-.25v2a1 1 0 0 1-1 1h-2m0-13h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m0 13h-2a1 1 0 0 1-1-1v-2"/></g>`), g.Group(children),
	)
}

func InterfaceUserRemoveActionsRemoveCloseGeometricHumanPersonMinusSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="3.75" r="2.25"/><path d="M6.5 13.5h-6V12a4.5 4.5 0 0 1 7.39-3.45m.61 2.95h5"/></g>`), g.Group(children),
	)
}

func InterfaceUserSingleCloseGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="3.75" r="3.25"/><path d="M13.18 13.5a6.49 6.49 0 0 0-12.36 0Z"/></g>`), g.Group(children),
	)
}

func InterfaceUserSquareAlternateSquareGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2.5"/><path d="M2.1 13.5a5 5 0 0 1 9.8 0"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`), g.Group(children),
	)
}

func InterfaceValidationCheckCheckFormValidationCheckmarkSuccessAddAddition(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 8.55l2.73 3.51a1 1 0 0 0 .78.39a1 1 0 0 0 .78-.36L13.5 1.55"/>`), g.Group(children),
	)
}

func InterfaceValidationCheckCircleCheckmarkAdditionCircleSuccessCheckValidationAddForm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4 8l2.05 1.64a.48.48 0 0 0 .4.1a.5.5 0 0 0 .34-.24L10 4"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func InterfaceValidationCheckSquareOneCheckFormValidationCheckmarkSuccessAddAdditionBoxSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="m4 8l2.05 1.64a.48.48 0 0 0 .4.1a.5.5 0 0 0 .34-.24L10 4"/></g>`), g.Group(children),
	)
}

func InterfaceValidationCheckSquareTwoCheckFormValidationCheckmarkSuccessAddAdditionBoxSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="m3 7l1.5 1.5L7 5m1 2.5h3"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherCelsiusDegreesTemperatureCentigradeCelsiusDegreeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><path d="M13.5 1.84a6 6 0 1 0-2 11.66a6 6 0 0 0 2-.34"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherCloudOneCloudMeteorologyCloudyOvercastCoverWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.726 5.706a2.34 2.34 0 0 0-.449 0a3.356 3.356 0 0 0-6.554 0a2.34 2.34 0 0 0-.45 0a2.614 2.614 0 1 0 0 5.219h7.453a2.613 2.613 0 1 0 0-5.219Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherCloudTwoCloudMeteorologyCloudyOvercastCoverWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 6.5a2.51 2.51 0 0 0-1.5.5h0A4.5 4.5 0 1 0 5 11.5h6a2.5 2.5 0 0 0 0-5Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherCresentMoonOneNightNewMoonCrescentWeatherTimeWaning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 7a7 7 0 0 1 3.44-6A6.58 6.58 0 0 0 9 .5a6.5 6.5 0 0 0 0 13a6.58 6.58 0 0 0 2.47-.5A7 7 0 0 1 8 7Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherCresentMoonTwoNightNewMoonCrescentWaxingTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 7a7 7 0 0 0-3.47-6A6.58 6.58 0 0 1 5 .5a6.5 6.5 0 1 1 0 13a6.58 6.58 0 0 1-2.47-.5A7 7 0 0 0 6 7Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherFahrenheitDegreesTemperatureFahrenheitDegreeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="2" r="1.5"/><path d="M7 13.5v-12h6M7 7h4"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherFirstQuarterMoonNightMoonFirstHalfQuarterTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.75.5a6.5 6.5 0 0 1 0 13Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherFullMoonNightFullMoonTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="7" cy="7" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`), g.Group(children),
	)
}

func InterfaceWeatherGibbousMoonOneNightMoonWeatherGibbousTimeWaning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7a6.5 6.5 0 0 0 6.5 6.5a6.58 6.58 0 0 0 1.19-.11a10 10 0 0 0 0-12.78A6.58 6.58 0 0 0 8.5.5A6.5 6.5 0 0 0 2 7Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherGibbousMoonTwoNightWaxingMoonGibbousTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 7a6.5 6.5 0 0 1-6.5 6.5a6.58 6.58 0 0 1-1.19-.11a10 10 0 0 1 0-12.78A6.58 6.58 0 0 1 5.5.5A6.5 6.5 0 0 1 12 7Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherHumidityNoneHumidityDropWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l13-13M4.85 13a4.5 4.5 0 0 0 6.65-4a5.48 5.48 0 0 0-.5-1.94M9.6 4.4C8.37 2.34 7 .5 7 .5S2.5 6.51 2.5 9a4.42 4.42 0 0 0 .5 2"/>`), g.Group(children),
	)
}

func InterfaceWeatherHurricaneTwisterTornadoHurricaneCycloneDisasterNaturalWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="3" rx="6" ry="2.5"/><path d="M11.87 6.73c-.87.79-2.73 1.34-4.94 1.34A10.17 10.17 0 0 1 3 7.37m7.86 2.47a6.84 6.84 0 0 1-3.36.73a8.47 8.47 0 0 1-2.5-.35m5.37 3.22c-.95.21-1.83-.16-2-.83"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherLightningOneCloudBoltStormWeatherThunderMeteorologyLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 7L6 10.5h3l-2 3"/><path d="M11 8a2.5 2.5 0 1 0-.62-4.92a3.5 3.5 0 0 0-6.76 0A2.5 2.5 0 1 0 3 8h1.5"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherLightningTwoCloudBoltStormWeatherThunderMeteorologyLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.78 9.38A2.5 2.5 0 1 0 9.5 5h0a4.5 4.5 0 1 0-5.75 4.32"/><path d="m8 8l-2 2.5h3l-2 3"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherMoonAstronomyMoonScienceSpaceCrescent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 10.48A6.55 6.55 0 0 1 6.46.5a6.55 6.55 0 0 0 1 13A6.46 6.46 0 0 0 13 10.39a6.79 6.79 0 0 1-1 .09Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherRainDropDropsRainRainyMeteorologyWaterPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9C11.5 6.51 7 .5 7 .5S2.5 6.51 2.5 9a4.5 4.5 0 0 0 9 0Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherRainOneCloudRainRainyMeteorologyPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.25 11.5l.5-1m4 1l.5-1M6 13.5l.5-1m-5 1l.5-1m8.5 1l.5-1M11 8a2.5 2.5 0 1 0-.62-4.92a3.5 3.5 0 0 0-6.76 0A2.5 2.5 0 1 0 3 8Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherRainTwoCloudRainRainyMeteorologyPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4 11.5l.5-1m4 1l.5-1m-3.25 3l.5-1m-5 1l.5-1m8.5 1l.5-1m-.5-4.5a2.5 2.5 0 0 0 0-5a2.54 2.54 0 0 0-1.57.55A3.75 3.75 0 1 0 5 8Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherSnowFlakeWinterFreezeSnowFreezingIceColdWeatherSnowflake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 .5l2 2l2-2M.5 9l2-2l-2-2M9 13.5l-2-2l-2 2M13.5 5l-2 2l2 2m-10-5.5L5 5m0 4l-1.5 1.5m7-7L9 5m0 4l1.5 1.5M7 2.5v9M2.5 7h9"/>`), g.Group(children),
	)
}

func InterfaceWeatherSnowOneCloudSnowSnowfallOvercastWeatherPrecipitationMeteorology(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.25 10a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm-2 2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm6-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM5 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm5.75-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM9 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm3.75 0a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 8a2.5 2.5 0 1 0-.62-4.92a3.5 3.5 0 0 0-6.76 0A2.5 2.5 0 1 0 3 8Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherSnowTwoCloudSnowSnowfallOvercastWeatherPrecipitationMeteorology(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.25 10a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm-2 2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm6-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM5 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm5.75-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM9 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm3.75 0a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8a2.5 2.5 0 0 0 0-5a2.54 2.54 0 0 0-1.57.55A3.75 3.75 0 1 0 5.25 8Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherSunOneCloudMeteorologyCloudyPartlySunnyWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 1.93a3 3 0 0 1 4.5 2.57a3.12 3.12 0 0 1-.21 1.11m-2.317 1.56a3.421 3.421 0 0 0-.46 0a3.43 3.43 0 0 0-6.71 0c-.15-.01-.3-.01-.45 0a2.67 2.67 0 1 0 0 5.33h7.62a2.67 2.67 0 0 0 0-5.33v0Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherSunPhotosLightCameraModeBrightnessSunPhotoFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="m13.5 7l-2.34 1.72l.44 2.88l-2.88-.44L7 13.5l-1.72-2.34l-2.88.44l.44-2.88L.5 7l2.34-1.72L2.4 2.4l2.88.44L7 .5l1.72 2.34l2.88-.44l-.44 2.88L13.5 7z"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherSunTwoCloudMeteorologyCloudyPartlySunnyWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11 8.5a2.51 2.51 0 0 0-1.5.5h0A4.5 4.5 0 1 0 5 13.5h6a2.5 2.5 0 0 0 0-5Z"/><circle cx="11.25" cy="2.75" r="2.25"/></g>`), g.Group(children),
	)
}

func InterfaceWeatherTemperatureColdTemperatureThermometerMinusMercuryColdWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9V2A1.5 1.5 0 0 0 10 .5h0A1.5 1.5 0 0 0 8.5 2v7a2.5 2.5 0 1 0 3 0Zm-10-6.5h3.25"/>`), g.Group(children),
	)
}

func InterfaceWeatherTemperatureHotTemperatureThermometerHotMercuryPlusWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9V2A1.5 1.5 0 0 0 10 .5h0A1.5 1.5 0 0 0 8.5 2v7a2.5 2.5 0 1 0 3 0Zm-10-6.5h3.98M3.49.51v3.98"/>`), g.Group(children),
	)
}

func InterfaceWeatherTemperatureTemperatureThermometerWeatherLevelMeterMercuryMeasure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9V2A1.5 1.5 0 0 0 7 .5h0A1.5 1.5 0 0 0 5.5 2v7a2.5 2.5 0 1 0 3 0Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherThirdQuarterMoonNightMoonThirdHalfQuarterTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25 13.5a6.5 6.5 0 0 1 0-13Z"/>`), g.Group(children),
	)
}

func InterfaceWeatherUmbrellaCloseStormRainUmbrellaCloseWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 9.5h-7l3-9h1l3 9zM7 9.5V12a1.5 1.5 0 0 0 3 0"/>`), g.Group(children),
	)
}

func InterfaceWeatherUmbrellaStormRainUmbrellaOpenWeather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7a5 5 0 0 1 10 0Zm5-5V.5M10 12a1.5 1.5 0 0 1-3 0V7"/>`), g.Group(children),
	)
}

func InterfaceWeatherWindWindHighOvercastGustWeatherMeteorologyGale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5.5a1.75 1.75 0 0 1 0 3.5h-7m11.25 6.5a1.75 1.75 0 0 0 0-3.5H2m5.25 6.5a1.75 1.75 0 0 0 0-3.5H1.5"/>`), g.Group(children),
	)
}

func InterfaceWeatherWindmillWindmillVelocityWeatherWind(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5h1a2 2 0 0 1 2 2V7h0h-3h0V.5h0ZM4 7h3v6.5h0h-1a2 2 0 0 1-2-2V7h0Zm9.5 0v1a2 2 0 0 1-2 2H7h0V7h6.5ZM7 4v3h0H.5h0V6a2 2 0 0 1 2-2H7Z"/>`), g.Group(children),
	)
}

func LegalJusticeHammerHammerWorkLegalMalletOfficeCompanyGavelJusticeJudgeArbitrationCourt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 13.29H8m-1 0v-2.5H1.5v2.5"/><rect width="7.07" height="4.24" x="3.96" y="2.17" rx="1" transform="rotate(-45 7.499 4.294)"/><path d="m9 5.79l4.5 4.5"/></g>`), g.Group(children),
	)
}

func LegalJusticeScaleOneOfficeWorkLegalScaleJusticeCompanyArbitrationBalanceCourt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.5L3 4L.5 9.5a2.5 2.5 0 0 0 5 0Zm8 0L11 4L8.5 9.5a2.5 2.5 0 0 0 5 0ZM1.5 4h11M7 4V2"/>`), g.Group(children),
	)
}

func LegalJusticeScaleTwoOfficeWorkLegalScaleJusticeUnequalCompanyArbitrationUnbalanceCourt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 11L3 4.82L.5 11a2.5 2.5 0 0 0 5 0Zm8-4L11 1.18L8.5 7a2.5 2.5 0 0 0 5 0Zm-12-1.5l11-5M7 3V.5"/>`), g.Group(children),
	)
}

func MailChatBubbleSquareMessagesMessageBubbleChatSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.5l-4 1l1-3v-9a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1Z"/>`), g.Group(children),
	)
}

func MailChatBubbleSquareWarningBubbleSquareMessagesNotificationChatMessageWarningAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.5 12.5l-4 1l1-3v-9a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1Zm3-9.5v3"/><circle cx="7.5" cy="9" r=".5"/></g>`), g.Group(children),
	)
}

func MailChatBubbleTypingOvalMessagesMessageBubbleTypingChat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.5" cy="7" r=".5"/><circle cx="6.75" cy="7" r=".5"/><circle cx="10" cy="7" r=".5"/><path d="M7 .5a6.5 6.5 0 0 0-5.41 10.1L.5 13.5l3.65-.66A6.5 6.5 0 1 0 7 .5Z"/></g>`), g.Group(children),
	)
}

func MailChatBubbleTypingSquareMessagesMessageBubbleTypingSquareChat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.25" cy="6.5" r=".5"/><circle cx="7.5" cy="6.5" r=".5"/><circle cx="10.75" cy="6.5" r=".5"/><path d="m4.5 12.5l-4 1l1-3v-9a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1Z"/></g>`), g.Group(children),
	)
}

func MailInboxEmailOutboxDrawerEmptyOpenInbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 8H4a1 1 0 0 1 1 1a2 2 0 0 0 4 0a1 1 0 0 1 1-1h3.5"/></g>`), g.Group(children),
	)
}

func MailSendEmailSendEmailPaperAirplane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.82 11L8 13.17a1.1 1.1 0 0 0 1.05.3a1.12 1.12 0 0 0 .81-.74L13.44 2A1.12 1.12 0 0 0 12 .56L1.27 4.14A1.12 1.12 0 0 0 .53 5a1.1 1.1 0 0 0 .3 1l2.74 2.74l-.09 3.47ZM13.12.78L3.57 8.74"/>`), g.Group(children),
	)
}

func MailSendEnvelopeEnvelopeEmailMessageUnopenedSealedClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10.5" x=".5" y="1.75" rx="1"/><path d="m.5 3l5.86 5a1 1 0 0 0 1.28 0l5.86-5"/></g>`), g.Group(children),
	)
}

func MailSendForwardEmailEmailSendMessageEnvelopeActionsActionForwardArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 12.5C3.42 10 4.14 9 7 9h1.5v3l5-5.5l-5-5v3h-1c-5 0-6 5-7 8Z"/>`), g.Group(children),
	)
}

func MailSendReplyAllEmailMessageReplyAllActionsActionArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 2.5l-3 3l3 3"/><path d="M13.5 11.5v-2a4 4 0 0 0-4-4h-5m-1-3l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func MailSendReplyEmailReplyMessageActionsActionArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 12.5C10.58 10 9.86 9 7 9H5.5v3l-5-5.5l5-5v3h1c5 0 6 5 7 8Z"/>`), g.Group(children),
	)
}

func MailSignAtEmailAtSignReadAddress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.05 7A3 3 0 1 1 7 4a3 3 0 0 1 3.05 3Z"/><path d="M10.05 7v1.3c0 3.49 5.47.2 2.6-4.54A6.59 6.59 0 0 0 7 .5A6.5 6.5 0 1 0 9.52 13"/></g>`), g.Group(children),
	)
}

func MailSignHashtagSharpSignHashtagTag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4.25h13m-13 5.5h13M11.25.5l-2.5 13m-3-13l-2.5 13"/>`), g.Group(children),
	)
}

func MailSmileyHappyFaceChatMessageSmileySmileEmojiFaceSatisfied(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M3.7 8c.5 1.8 2.5 2.9 4.3 2.4c1.1-.4 2-1.3 2.3-2.4M4.8 5.45a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5m4.4.5a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5"/></g>`), g.Group(children),
	)
}

func MailSmileyHappyMaskChatMessageSmileySmileEmojiMask(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 8s.5 1.5 3 1.5S10 8 10 8"/><path d="M13.5 7a6.5 6.5 0 0 1-13 0V.5h13Z"/><path d="M5 5a1 1 0 0 0-2 0m8 0a1 1 0 0 0-2 0"/></g>`), g.Group(children),
	)
}

func MailSmileySadFaceChatMessageSmileyEmojiSadFaceUnsatisfied(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M3.7 10.5C4.2 8.7 6.1 7.6 8 8.1c1.1.3 2 1.2 2.4 2.4M4.85 5.45a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5m4.4.5a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5"/></g>`), g.Group(children),
	)
}

func MailSmileySadMaskChatMessageSmileyEmojiSadMaskCry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7a6.5 6.5 0 0 1-13 0V.5h13Z"/><path d="M4 9.5S4.5 8 7 8s3 1.5 3 1.5m-5-5a1 1 0 0 1-2 0m8 0a1 1 0 0 1-2 0"/></g>`), g.Group(children),
	)
}

func MailSmileyStraightFaceChatMessageIndifferentSmileyEmojiFacePoker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 9.5h7m-3.5 4a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M4.8 5.45a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5m4.4.5a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5"/></g>`), g.Group(children),
	)
}

func MoneyAtmCardOneCreditPayPaymentDebitCardFinancePlasticMoney(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9.5" x=".5" y="2.25" rx="1"/><path d="M.5 5.75h13m-4 3.5H11"/></g>`), g.Group(children),
	)
}

func MoneyAtmCardThreeDepositMoneyPaymentFinanceAtmWithdraw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y="6.5" rx="1"/><path d="M3.5 2v2M7 .5V4m3.5-2v2"/><circle cx="7" cy="10" r="1.5"/></g>`), g.Group(children),
	)
}

func MoneyAtmCardTwoDepositMoneyPaymentFinanceAtmWithdraw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1 4.5A.5.5 0 0 1 .5 4V1A.5.5 0 0 1 1 .5h12a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5"/><rect width="7" height="8" x="3.5" y="3" rx=".5"/><circle cx="7" cy="7" r="1.5"/><path d="M3.5 13.5h7"/></g>`), g.Group(children),
	)
}

func MoneyBankInstitutionMoneySavingBankPaymentFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.91 5.5H1.09c-.56 0-.8-.61-.36-.9L6.64.73a.71.71 0 0 1 .72 0l5.91 3.87c.44.29.2.9-.36.9Z"/><rect width="13" height="2.5" x=".5" y="11" rx=".5"/><path d="M2 5.5V11m2.5-5.5V11M7 5.5V11m2.5-5.5V11M12 5.5V11"/></g>`), g.Group(children),
	)
}

func MoneyCashBagBagPaymentCashMoneyFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 3.02L13.5.5M8 3.02l5.5 1.48m-7 9c3.5 0 6-1.24 6-4c0-3-1.5-5-4.5-6.5l1.18-1.53a.65.65 0 0 0-.56-.95H4.38a.65.65 0 0 0-.56 1L5 3C2 4.52.5 6.52.5 9.52c0 2.74 2.5 3.98 6 3.98Z"/>`), g.Group(children),
	)
}

func MoneyCashBagDollarBagPaymentCashMoneyFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 10.02v1.01m0-6.02v.94m0 7.54c3.5 0 6-1.24 6-4c0-3-1.5-5-4.5-6.5l1.18-1.52a.66.66 0 0 0-.56-1H4.88a.66.66 0 0 0-.56 1L5.5 3C2.5 4.51 1 6.51 1 9.51c0 2.74 2.5 3.98 6 3.98Z"/><path d="M6 9.56A1.24 1.24 0 0 0 7 10a1.12 1.12 0 0 0 1.19-1A1.12 1.12 0 0 0 7 8a1.12 1.12 0 0 1-1.19-1A1.11 1.11 0 0 1 7 6a1.26 1.26 0 0 1 1 .4"/></g>`), g.Group(children),
	)
}

func MoneyCashBillOneBillingBillsPaymentFinanceCashCurrencyMoneyAccounting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y=".5" rx="1"/><circle cx="7" cy="4.5" r="1.5"/><path d="M1.5 11h11m-10 2.5h9"/></g>`), g.Group(children),
	)
}

func MoneyCashBillThreeAccountingBillingPaymentFinanceCashCurrencyMoneyBill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="8" x=".5" y="1.75" rx="1"/><circle cx="5.75" cy="5.75" r="1.5"/><path d="M3.5 12.25h9a1 1 0 0 0 1-1v-5"/></g>`), g.Group(children),
	)
}

func MoneyCashBillTwoCurrencyBillingPaymentFinanceCashBillMoneyAccounting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx="1"/><circle cx="7" cy="7" r="1.5"/><path d="M3 5h.5m7 4h.5"/></g>`), g.Group(children),
	)
}

func MoneyCashCoinsStackAccountingBillingPaymentStackCashCoinsCurrencyMoneyFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="9" cy="5.5" rx="4.5" ry="2"/><path d="M4.5 5.5v6c0 1.1 2 2 4.5 2s4.5-.9 4.5-2v-6"/><path d="M13.5 8.5c0 1.1-2 2-4.5 2s-4.5-.9-4.5-2m4.4-7A6.77 6.77 0 0 0 5 .5C2.51.5.5 1.4.5 2.5c0 .59.58 1.12 1.5 1.5"/><path d="M2 10C1.08 9.62.5 9.09.5 8.5v-6"/><path d="M2 7C1.08 6.62.5 6.09.5 5.5"/></g>`), g.Group(children),
	)
}

func MoneyCashDollarCoinAccountingBillingPaymentCashCoinCurrencyMoneyFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4.5V3M5.5 8.5c0 .75.67 1 1.5 1s1.5 0 1.5-1c0-1.5-3-1.5-3-3c0-1 .67-1 1.5-1s1.5.38 1.5 1M7 9.5V11"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func MoneyCashFileDollarCommonMoneyCurrencyCashFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5.5h-7a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-7Z"/><path d="M8.5 5V.5l5 5H9a.5.5 0 0 1-.5-.5Zm-4-.5V3M3 8.5c0 .75.67 1 1.5 1s1.5 0 1.5-1C6 7 3 7 3 5.5c0-1 .67-1 1.5-1s1.5.38 1.5 1m-1.5 4V11m4-1.5h3"/></g>`), g.Group(children),
	)
}

func MoneyCashSearchDollarSearchPayProductCurrencyQueryMagnifyingCashBusinessMoneyGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.5" cy="6.5" r="6"/><path d="m10.74 10.74l2.76 2.76M6.5 4V2.5M5 8c0 .75.67 1 1.5 1S8 9 8 8c0-1.5-3-1.5-3-3c0-1 .67-1 1.5-1S8 4.38 8 5M6.5 9v1.5"/></g>`), g.Group(children),
	)
}

func MoneyCashierBarCodeCodesTagsUpcBarcode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5v10m3-10v10m3-10v10M11 .5v10m2.5-10v10m-13 3h13"/>`), g.Group(children),
	)
}

func MoneyCashierQrCodeCodesTagsCodeQr(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4" height="4" x="3" y="3" rx="1"/><path d="M3 .5H1.5a1 1 0 0 0-1 1V3M11 .5h1.5a1 1 0 0 1 1 1V3M3 13.5H1.5a1 1 0 0 1-1-1V11M11 13.5h1.5a1 1 0 0 0 1-1V11M3 9.5V11h1.5M7 11V9.5H5.5m5.5-5H9.5V3M11 8V6.5H9.5m0 3V11H11"/></g>`), g.Group(children),
	)
}

func MoneyCashierShopShoppingPayPaymentCashierStoreCashRegisterMachine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 11h1M10 5V3m2-1.25A1.25 1.25 0 0 1 10.75 3h-1.5a1.25 1.25 0 0 1 0-2.5h1.5A1.25 1.25 0 0 1 12 1.75ZM5.5 5V1.5h-3V5"/><rect width="13" height="5" x=".5" y="8.5" rx="1"/><path d="M12.5 8.5V6a1 1 0 0 0-1-1h-9a1 1 0 0 0-1 1v2.5"/></g>`), g.Group(children),
	)
}

func MoneyCashierTagCodesTagsTagProductLabel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.25 9.36l-3.89 3.89a.75.75 0 0 1-1.06 0L1.79 6.74a.36.36 0 0 1-.11-.29l.59-3.84a.37.37 0 0 1 .34-.34l3.84-.59a.36.36 0 0 1 .29.11l6.51 6.51a.75.75 0 0 1 0 1.06ZM4.03 4.03L.53.53"/>`), g.Group(children),
	)
}

func MoneyCurrencyBitcoinCircleOneCryptoCirclePaymentBlockchainFinanceBitcoinCurrencyMoney(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.88 6.66h0a1.37 1.37 0 0 0 1.39-1.37h0A1.37 1.37 0 0 0 7.9 3.92H5.79a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h2.09a1.62 1.62 0 1 0 0-3.24Zm.02 0H5.29m.95-2.74V3m1.54.92V3m-1.54 8V9.9M7.78 11V9.9"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func MoneyCurrencyBitcoinCryptoCirclePaymentBlokchainFinanceBitcoinCurrencyMoney(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.79 6.45h0A2.23 2.23 0 0 0 7.83 2H4.39a.81.81 0 0 0-.81.81v8.09a.81.81 0 0 0 .81.81h3.4a2.63 2.63 0 0 0 0-5.26ZM5.13 2V.5M7.63 2V.5m-2.5 13V12m2.5 1.5V12m.2-5.55H3.58"/>`), g.Group(children),
	)
}

func MoneyCurrencyDollarDollarExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.5h3.85a1.65 1.65 0 0 0 .32-3.27l-3.34-.46a1.65 1.65 0 0 1 .32-3.27H9.5M7 13.5V.5"/>`), g.Group(children),
	)
}

func MoneyCurrencyEuroCircleExchangePaymentEuroForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 9.8a2.59 2.59 0 0 1-1 .2a2.88 2.88 0 0 1-2.75-3A2.88 2.88 0 0 1 7.5 4a2.76 2.76 0 0 1 .82.13M3.5 6h3m-3 2h3"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func MoneyCurrencyEuroExchangePaymentEuroForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 12.59a5 5 0 0 1-2 .41a5.77 5.77 0 0 1-5.5-6A5.77 5.77 0 0 1 10 1a4.89 4.89 0 0 1 1.63.27M2 5.5h6m-6 3h6"/>`), g.Group(children),
	)
}

func MoneyCurrencyPoundCirclePoundSterlingExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.06 5.42v-.36A1.56 1.56 0 0 0 7.5 3.5h0a1.56 1.56 0 0 0-1.56 1.56v3L5 9.5h4.06M4.5 6.62h2.88"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func MoneyCurrencyPoundPoundSterlingExchangePaymentForexFinanceCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.75 4.75V4A3.26 3.26 0 0 0 8.5.75h0A3.26 3.26 0 0 0 5.25 4v6.25l-2 3h8.5m-9.5-6h5.99"/>`), g.Group(children),
	)
}

func MoneyCurrencyYuanCircleExchangePaymentForexFinanceYuanCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m5 4l2.08 2.77L9.15 4M7.08 6.77V10m-1.62-.69h3.23M5.46 7.5h3.23"/><circle cx="7" cy="7" r="6.5"/></g>`), g.Group(children),
	)
}

func MoneyCurrencyYuanExchangePaymentForexFinanceYuanCurrencyMoneyForeign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5.5l4.5 6l4.5-6M7 6.5v7m-3.5-3h7m-7-2h7"/>`), g.Group(children),
	)
}

func MoneyGraphAnalyticsBusinessProductGraphDataChartAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13h13"/><path d="M3.5 6.5L6 9l4-6l3.5 2.5"/></g>`), g.Group(children),
	)
}

func MoneyGraphArrowDecreaseDownStatsGraphDescendRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 10.5h4v-4"/><path d="M13.5 10.5L7.85 4.85a.5.5 0 0 0-.7 0l-2.3 2.3a.5.5 0 0 1-.7 0L.5 3.5"/></g>`), g.Group(children),
	)
}

func MoneyGraphArrowIncreaseAscendGrowthUpArrowStatsGraphRightGrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 3.5h4v4"/><path d="M13.5 3.5L7.85 9.15a.5.5 0 0 1-.7 0l-2.3-2.3a.5.5 0 0 0-.7 0L.5 10.5"/></g>`), g.Group(children),
	)
}

func MoneyGraphBarDecreaseArrowProductPerformanceDownDecreaseGraphBusinessChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.24.5l11.5 5.23m-2.15.81l2.15-.81l-.8-2.16M1.25 6h1.5a.5.5 0 0 1 .5.5v7h0h-2.5h0v-7a.5.5 0 0 1 .5-.5Zm5 1.5h1.5a.5.5 0 0 1 .5.5v5.5h0h-2.5h0V8a.5.5 0 0 1 .5-.5Zm5 1.5h1.5a.5.5 0 0 1 .5.5v4h0h-2.5h0v-4a.5.5 0 0 1 .5-.5Z"/>`), g.Group(children),
	)
}

func MoneyGraphBarIncreaseUpProductPerformanceIncreaseArrowGraphBusinessChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.24 6.54l11.5-5.23M10.59.5l2.15.81l-.8 2.15m1.31 10.05h-2.5h0v-7a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5v7h0Zm-5 0h-2.5h0v-5.5a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5v5.5h0Zm-5 0H.75h0v-4a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 .5.5v4h0Z"/>`), g.Group(children),
	)
}

func MoneyGraphBarProductDataBarsAnalysisAnalyticsGraphBusinessChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5h13m-9 0V.5h-4v13m8 0v-7h-4v7m8 0v-10h-4v10"/>`), g.Group(children),
	)
}

func MoneyGraphPieChartProductDataAnalysisAnalyticsPieGraphBusinessChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5V7l4.6 4.6"/></g>`), g.Group(children),
	)
}

func MoneySafeVaultSavingComboPaymentSafeMoneyCombinationFinance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="11.5" x=".5" y=".5" rx="1"/><circle cx="8.5" cy="6.25" r="1.75"/><path d="M8.5 3.25V4.5m0 3.5v1.25m3-3h-1.25m-3.5 0H5.5m5.12-2.12l-.88.88M7.26 7.49l-.88.88m4.24 0l-.88-.88M7.26 5.01l-.88-.88M3 4.5V8m-1 4v1.5m9.5-1.5v1.5"/></g>`), g.Group(children),
	)
}

func MoneyWalletMoneyPaymentFinanceWallet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 7.5v-2a1 1 0 0 0-1-1H1.5a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1H11a1 1 0 0 0 1-1V10M3.84 2L9.51.52a.49.49 0 0 1 .61.36L10.4 2"/><rect width="3.5" height="2.5" x="10" y="7.5" rx=".5"/></g>`), g.Group(children),
	)
}

func NatrueEcologyRecycleOneSignEnvironmentProtectSaveArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6 10.08v2a.29.29 0 0 1-.29.28h-2.1a1.45 1.45 0 0 1-.88-.51"/><path d="M2.73 11.86s-1.09-.95 1.44-3.48H5.3L3.61 5.59H.5l.85 1.09s-.57.1-.57.89a10 10 0 0 0 1.95 4.29Zm7.62-5.49l1.48-1a.26.26 0 0 1 .21 0a.26.26 0 0 1 .19.12l1.15 1.72a1 1 0 0 1 0 .89"/><path d="M13.42 8s-.32 1.3-3.63 1v-.87l-1.94 2.54l1.84 2.83l.1-1.33s.37.45 1 0A18.73 18.73 0 0 0 13.42 8ZM4.35 4.28L3 3.66a.3.3 0 0 1-.14-.38L4 1a1 1 0 0 1 .74-.5m0 0s1.23 0 2.79 3l-.81 1h3.36l1.19-3.14l-1.4.43a1.21 1.21 0 0 0-.73-1A5.87 5.87 0 0 0 7 .5Z"/></g>`), g.Group(children),
	)
}

func NatureEcologyBonePetDogBoneFoodSnack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 2.5a2 2 0 0 0-4 0a2 2 0 0 0 .59 1.41L3.91 8.09A2 2 0 0 0 2.5 7.5a2 2 0 0 0 0 4a2 2 0 0 0 4 0a2 2 0 0 0-.59-1.41l4.18-4.18a2 2 0 0 0 1.41.59a2 2 0 0 0 0-4Z"/>`), g.Group(children),
	)
}

func NatureEcologyCatHeadCatPetAnimalsFelyne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 10.25h1m-7 3l3.5-2m-3.5-3l4 1m2.98-6.97L12.87.77a.46.46 0 0 1 .46.11a.49.49 0 0 1 .16.45l-1.1 5.61c-.06-.15-.12-.31-.17-.47A5.75 5.75 0 0 0 7 2.25a5.75 5.75 0 0 0-5.22 4.22c-.05.17-.11.32-.17.48L.51 1.33A.49.49 0 0 1 .67.88a.46.46 0 0 1 .46-.11l5.39 1.51m1.98 10.8a6.7 6.7 0 0 1-3 0m8 .17l-3.5-2m3.5-3l-4 1"/>`), g.Group(children),
	)
}

func NatureEcologyCloverPlantLeafTreeFlowerLuckLucky(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 7s2-.5 2-2c0-1-.5-1.5-2-1.5c-1 0-2.08 1.83-2.5 2c.17-.42 1.5-2 1.5-3c0-1.5-.5-2-1.5-2c-1.5 0-2 2-2 2s-.5-2-2-2c-1 0-1.5.5-1.5 2c0 1 1.33 2.08 1.5 2.5c-.42-.17-1.5-1.5-2.5-1.5C1 3.5.5 4 .5 5c0 1.5 2 2 2 2s-2 .5-2 2c0 1 .5 1.5 2 1.5a6.36 6.36 0 0 0 4.5-2a6.36 6.36 0 0 0 4.5 2c1.5 0 2-.5 2-1.5c0-1.5-2-2-2-2ZM7 8.5v5"/>`), g.Group(children),
	)
}

func NatureEcologyDogHeadDogPetAnimalsCanine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 5v2a1 1 0 0 0 1 1h.5a2 2 0 0 0 2-2V3a1 1 0 0 0-1-1H11a2 2 0 0 1-.91-.22A6.88 6.88 0 0 0 7 1a6.88 6.88 0 0 0-3.11.78A2.07 2.07 0 0 1 3 2H1.5a1 1 0 0 0-1 1v3a2 2 0 0 0 2 2H3a1 1 0 0 0 1-1V5m7.5 3l-.31 1.57a4.27 4.27 0 0 1-8.38 0L2.5 8m4 2h1"/>`), g.Group(children),
	)
}

func NatureEcologyFlowerPlantTreeFlowerPetalsBloom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="7.26" rx="1.5" ry="1.49"/><path d="M13.37 5.34a2.57 2.57 0 0 0-3.24-1.64a2.3 2.3 0 0 0-.65.3a2.71 2.71 0 0 0 .1-.71a2.58 2.58 0 0 0-5.16 0a2.71 2.71 0 0 0 .1.71a2.3 2.3 0 0 0-.65-.31a2.56 2.56 0 1 0-1.59 4.87a3 3 0 0 0 .72.12a3 3 0 0 0-.5.51a2.54 2.54 0 0 0 .57 3.57a2.59 2.59 0 0 0 3.6-.56a2.47 2.47 0 0 0 .33-.64a2.47 2.47 0 0 0 .34.64a2.59 2.59 0 0 0 3.6.56a2.54 2.54 0 0 0 .57-3.57a3 3 0 0 0-.5-.51a3 3 0 0 0 .71-.12a2.55 2.55 0 0 0 1.65-3.22Z"/></g>`), g.Group(children),
	)
}

func NatureEcologyGreenHouseGlassBuildingPlantsCropsProduceFarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.94a1 1 0 0 0-.32-.74L7 .5L.82 6.2a1 1 0 0 0-.32.74v5.56a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1ZM7 8V.5M.5 8h13m-13 2.75H4m6 0h3.5M4 13.5V3.27m6 10.23V3.27"/>`), g.Group(children),
	)
}

func NatureEcologyLeafEnvironmentLeafEcologyPlantPlantsEco(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.51 5.38c2 2.82.47 6.05-.27 7.31a1.42 1.42 0 0 1-1 .66c-1.45.25-5.06.53-7-2.29C1.33 8.4 1.41 3.72 1.58 1.49A1.05 1.05 0 0 1 3 .55c2.15.62 6.63 2.17 8.51 4.83Z"/><path d="M4.77 4.45a52.26 52.26 0 0 1 6 8.73"/></g>`), g.Group(children),
	)
}

func NatureEcologyLeafProtectEnvironmentLeafEcologyPlantPlantsEco(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.31 4.15c0-1 0-2-.07-2.66A1 1 0 0 0 9.89.55c-2.16.62-6.57 2.17-8.42 4.83c-2 2.82-.46 6.05.27 7.31a1.38 1.38 0 0 0 1 .66a9.36 9.36 0 0 0 2.79.08"/><path d="M5 8.58a51.12 51.12 0 0 0-2.78 4.6m8.45.29h0a.5.5 0 0 1-.34 0h0A4.48 4.48 0 0 1 7.5 9.31V7.46A.47.47 0 0 1 8 7h5a.47.47 0 0 1 .46.46v1.85a4.48 4.48 0 0 1-2.79 4.16Z"/></g>`), g.Group(children),
	)
}

func NatureEcologyPineTreePlantTreeFarmingChristmasNaturePlantsPineEnvironment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 10.5H2l3.5-5H3l4-5l4 5H8.5l3.5 5zm-5 0v3"/>`), g.Group(children),
	)
}

func NatureEcologyPottedCactusTreePlantSucculentPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.1 12.61a1 1 0 0 1-1 .89H4.9a1 1 0 0 1-1-.89L3.5 9h7ZM4.5 9V5a2.5 2.5 0 0 1 5 0v4M7 2.5v-2M4.5 6h-2m1-3.5l1.31 1.31M9.5 6h2m-1-3.5L9.19 3.81"/>`), g.Group(children),
	)
}

func NatureEcologyPottedFlowerFlowerPlantTreePot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 13.5H5l-1-4h6l-1 4zm1.31-10.61A1.34 1.34 0 0 0 8.63 2a1.2 1.2 0 0 0-.34.17a1.5 1.5 0 0 0 .05-.37a1.34 1.34 0 0 0-2.68 0a1.5 1.5 0 0 0 0 .37A1.2 1.2 0 0 0 5.37 2a1.34 1.34 0 0 0-1.68.86a1.32 1.32 0 0 0 .86 1.67a1.15 1.15 0 0 0 .37.06A1.34 1.34 0 0 0 5 6.75a1.34 1.34 0 0 0 1.87-.3A1.06 1.06 0 0 0 7 6.12a1.06 1.06 0 0 0 .18.33a1.34 1.34 0 0 0 1.87.3a1.34 1.34 0 0 0 0-2.13a1.15 1.15 0 0 0 .37-.06a1.32 1.32 0 0 0 .89-1.67ZM7 6.12V9.5"/>`), g.Group(children),
	)
}

func NatureEcologyPottedPlantTreePlantSucculentPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 7.5l-.86 5.16a1 1 0 0 1-1 .84H4.35a1 1 0 0 1-1-.84L2.5 7.5ZM4 .69a3.84 3.84 0 0 0-1.5 3.06A3.63 3.63 0 0 0 6 7.5a3.24 3.24 0 0 0 .94-.14Zm6 0a3.84 3.84 0 0 1 1.5 3.06A3.63 3.63 0 0 1 8 7.5a3.24 3.24 0 0 1-.94-.14Z"/><path d="M5 3L7 .5L9 3"/></g>`), g.Group(children),
	)
}

func NatureEcologyPottedTreeOneTreePlantPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="4" r="3.5"/><path d="M7 5v4.5m2 4H5l-.5-4h5l-.5 4z"/></g>`), g.Group(children),
	)
}

func NatureEcologyPottedTreeTwoTreePlantPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.25 13.5h-4l-.5-4h5l-.5 4zm2.88-7.43a2 2 0 1 1-2.47-3.14c.86-.68 3.59-.28 3.59-.28S13 5.39 12.13 6.07Z"/><path d="M9.25 5.64a5.5 5.5 0 0 0-2 3.86c0-3-1.5-5-4-6.5"/><path d="M1.67 5a2.56 2.56 0 0 0 3.61-3.61C4.28.4.75.5.75.5S.67 4 1.67 5Z"/></g>`), g.Group(children),
	)
}

func NatureEcologyRainbowArchRainColorfulRainbowCurveHalfCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.25a6.5 6.5 0 0 1 13 0"/><path d="M3 10.25a4 4 0 0 1 8 0"/><path d="M5.5 10.25a1.5 1.5 0 0 1 3 0"/></g>`), g.Group(children),
	)
}

func NatureEcologyRiceFieldSunRiseSetFieldCropProduceFarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 11.5s2-2 6.5-2s6.5 2 6.5 2M7 9.5v4m-4.5 0L6 9.53m5.5 3.97L8 9.53M3.54 7.5a2.74 2.74 0 0 1 0-.5a3.5 3.5 0 0 1 7 0a2.74 2.74 0 0 1 0 .5M.5 7h1m1-4.5L3 3M7 .5v1m4.5 1L11 3m2.5 4h-1"/>`), g.Group(children),
	)
}

func NatureEcologyRoseFlowerRosePlantTree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4.1a8.62 8.62 0 0 1 5.5-1.6V6a5.5 5.5 0 0 1-11 0V2.5s6.56-.79 9.56 7.21"/><path d="M9.91 2.76a3 3 0 0 0-5.82 0M7 11.5v2"/></g>`), g.Group(children),
	)
}

func NatureEcologyTreeOneTreePlantPineTrianglePark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.57 7.84a2 2 0 0 1-.26 1.82a2 2 0 0 1-1.63.84H4.32a2 2 0 0 1-1.63-.84a2 2 0 0 1-.26-1.82l1.91-5.45a2.82 2.82 0 0 1 5.32 0Z"/><path d="M5.5 6.5L7 8v5.5M7 8l1.5-1.5"/></g>`), g.Group(children),
	)
}

func NatureEcologyTreeTwoTreePlantCircleRoundPark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 6.5L7 8v5.5M7 8l1.5-1.5"/><circle cx="7" cy="6" r="5.5"/></g>`), g.Group(children),
	)
}

func NatureEcologyVolcanoEruptionEruptMountainVolcanoLavaMagmaExplosion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5s-4-2-4-5h-5c0 3-4 5-4 5m10-8a2 2 0 0 0 0-4a2 2 0 0 0-1.24.43a2.5 2.5 0 0 0-4.52 0A2 2 0 0 0 3.5 1.5a2 2 0 0 0 0 4m2-1v4m3-4v4"/>`), g.Group(children),
	)
}

func PhoneContactPhonebookPhonebookPhoneNumberBooksBook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 13.5H2a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1ZM4 .5v13M7.5 4h2"/>`), g.Group(children),
	)
}

func PhoneMobilePhoneAndroidPhoneMobileDeviceSmartphoneIphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="13" x="2" y=".5" rx="1"/><path d="M6.25 11h1.5"/></g>`), g.Group(children),
	)
}

func PhoneModeAirplaneServerPlaneAirplaneDisableWirelessModeInternetNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.39 10.61H1.94a1.45 1.45 0 0 1 0-2.89h2.3l1.48-1.48l-4.06-2.5a1.45 1.45 0 0 1-.52-1.91a1.47 1.47 0 0 1 2-.77l5.3 2.45L11 .92A1.45 1.45 0 0 1 13.08 3l-2.59 2.56l2.45 5.35a1.46 1.46 0 0 1-.77 1.95a1.45 1.45 0 0 1-1.91-.52l-2.5-4.06l-1.48 1.48v2.3a1.45 1.45 0 0 1-2.89 0Z"/>`), g.Group(children),
	)
}

func PhoneSignalFullPhoneMobileDeviceSignalWirelessSmartphoneIphoneBarBarsFullAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75 7.5H.5v6h4.25m4.5-9h-4.5v9h4.5m0-13h4.25v13H9.25z"/>`), g.Group(children),
	)
}

func PhoneSignalLowPhoneMobileDeviceSignalWirelessSmartphoneIphoneBarLowBarsAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7.5h4.25v6H.5zm10.5 6h2.5m-8.75 0H8.5"/>`), g.Group(children),
	)
}

func PhoneSignalMediumSmartphonePhoneMobileDeviceIphoneSignalMediumWirelessBarBarsAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75 7.5H.5v6h4.25m0-9h4.5v9h-4.5zm4.5 9h4.25"/>`), g.Group(children),
	)
}

func PhoneSignalNonePhoneMobileDeviceSignalWirelessSmartphoneIphoneBarBarsNoZeroAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5H3m8 0h2.5m-8 0h3"/>`), g.Group(children),
	)
}

func PhoneTelephoneAndroidPhoneMobileDeviceSmartphoneIphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.76 13a3.19 3.19 0 0 0 4-.44l.45-.44a1.08 1.08 0 0 0 0-1.51L11.3 8.72a1.07 1.07 0 0 0-1.5 0h0a1.08 1.08 0 0 1-1.51 0l-3-3a1.06 1.06 0 0 1 0-1.51h0a1.07 1.07 0 0 0 0-1.5L3.39.81a1.08 1.08 0 0 0-1.51 0l-.44.45a3.19 3.19 0 0 0-.44 4A28.94 28.94 0 0 0 8.76 13Z"/>`), g.Group(children),
	)
}

func ProgrammingBrowserAddAppCodeAppsAddProgrammingWindowPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-4 5h-5M7 6v5"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserBuildBuildWebsiteDevelopmentWindowCodeProgrammingWebBackendBrowserDev(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7.5v5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h4v7Zm-13-4h13m-3.5 4v-7"/>`), g.Group(children),
	)
}

func ProgrammingBrowserCheckCheckmarkPassWindowAppCodeProgrammingSuccessCheckApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13M4 9l2 1.5l3.5-4"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserCodeOneCodeBrowserLineShellProgrammingCommandTerminal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 4h13M4 7l1.5 1.5L4 10m4.5-1.5h2"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserCodeTwoCodeBrowserTagsAngleProgrammingBracket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 4h13m-9 3L3 8.5L4.5 10M10 7l1.5 1.5L10 10m-3.5.5L8 6"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserDashboardAppCodeAppsProgrammingWindowDashboardPerformance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-6.5 7L10.5 7m0 3.5h1M7 7V6M4.53 8.03l-.71-.71M3.5 10.5h-1"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserDeleteAppCodeAppsFailDeleteWindowRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-4.5 7l-4-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserFavoriteHeartWindowAppCodeFavoriteLikeProgrammingHeartApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13M7 7.5c1-2 3-1 3 .5c0 2-3 3-3 3s-3-1-3-3c0-1.5 2-2.5 3-.5Z"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserFavoriteStarWindowStarAppCodeFavoriteLikeProgrammingApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 13.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1m-12-10h13"/><path d="M6.45 7.37a.59.59 0 0 1 1.1 0L8.19 9H9.9a.6.6 0 0 1 .4 1l-1.51 1.36l.64 1.28a.58.58 0 0 1-.12.69a.61.61 0 0 1-.7.1L7 12.55l-1.61.88a.61.61 0 0 1-.7-.1a.58.58 0 0 1-.12-.69l.64-1.28L3.7 10a.6.6 0 0 1 .4-1h1.71Z"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserHashWindowHashCodeProgrammingInternetLanguageBrowserWebTag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 13.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1m-12-10h13m-7 3l-2 7m5-7l-2 7m3-5H4m6 3H3.5"/>`), g.Group(children),
	)
}

func ProgrammingBrowserKeySecurePasswordWindowBrowserKeySecurityLogin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.44" cy="11.33" r="2.17"/><path d="m8 9.8l3.86-3.86a.36.36 0 0 1 .51 0l1.13 1.15m-3.05.28l1.02 1.02M2 12.5h-.5a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1V4m-13-.5h13"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserMultipleWindowAppCodeAppsTwoProgrammingWindowCascade(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 5V1.5a1 1 0 0 1 1-1h8.5a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H11m-8-7h10.5"/><rect width="8" height="6" x=".5" y="7.5" rx="1"/><path d="M.5 10h8"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserRemoveAppCodeAppsSubtractProgrammingWindowMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-4 5h-5"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserSearchSearchWindowGlassAppCodeProgrammingQueryFindMagnifyingApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 12.5h-4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v5m-13-3h13"/><circle cx="9.5" cy="9.5" r="2.5"/><path d="m11.27 11.27l2.23 2.23"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserSettingWindowGearAppCodeProgrammingCogSettingsApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 13.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1m-12-10h13M7 6.5V8m-3.03.25l1.3.75m-1.3 2.75l1.3-.75M7 13.5V12m3.03-.25L8.73 11m1.3-2.75L8.73 9"/><circle cx="7" cy="10" r="2"/></g>`), g.Group(children),
	)
}

func ProgrammingBrowserWarningWindowAppCodeProgrammingNotificationWarningAlertApps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 13.5h-2a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-2m-10-10h13M7 6.5V10"/><circle cx="7" cy="13" r=".5"/></g>`), g.Group(children),
	)
}

func ProgrammingBugCodeBugSecurityProgrammingSecureComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="8" r="4.5"/><path d="M2.5 8h-2m0 3.5a3.46 3.46 0 0 0 2.63-1.2m0-4.6A3.46 3.46 0 0 0 .5 4.5m11 3.5h2m0 3.5a3.46 3.46 0 0 1-2.63-1.2m0-4.6a3.46 3.46 0 0 1 2.63-1.2m-5.26-.83A2.5 2.5 0 0 0 9.5 1.5m-5 0a2.5 2.5 0 0 0 1.26 2.17M2.61 7h8.78"/></g>`), g.Group(children),
	)
}

func ProgrammingCloudAddCloudNetworkInternetAddServerPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 8.5M10 11H5m2.5-2.5v5"/>`), g.Group(children),
	)
}

func ProgrammingCloudCloudInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 11a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 11Z"/>`), g.Group(children),
	)
}

func ProgrammingCloudDataTransferCloudDataTransferInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 8.1a3 3 0 1 0-3.65-4.69A4 4 0 1 0 1.38 7M3.5 9.5l2-2v6"/><path d="m10.5 11.5l-2 2v-6"/></g>`), g.Group(children),
	)
}

func ProgrammingCloudDownloadCloudDownInternetNetworkDownloadServerArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 8.5m3 5v-6m-2 4l2 2l2-2"/>`), g.Group(children),
	)
}

func ProgrammingCloudOffCloudNetworkInternetDisableServerOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.17 3.17a4.08 4.08 0 0 1 6.18 2.24a3 3 0 1 1 2.24 5.09l-1 .05M.75 5.1A4 4 0 0 0 .5 6.5a4 4 0 0 0 4 4H6m-5.5-9l11 11"/>`), g.Group(children),
	)
}

func ProgrammingCloudRemoveCloudNetworkInternetSubtractMinusServerRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 8.5Zm-.5 3H5"/>`), g.Group(children),
	)
}

func ProgrammingCloudUploadCloudServerInternetUploadUpArrowNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7.5v6m-2-4l2-2l2 2"/><path d="M12 8.1a3 3 0 1 0-3.65-4.69A4 4 0 0 0 .5 4.5A4 4 0 0 0 1.38 7"/></g>`), g.Group(children),
	)
}

func ProgrammingCloudWifiCloudWifiInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 13.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 13.5Zm-.02-12.3a9 9 0 0 0-7 0m5.7 2.16a6.78 6.78 0 0 0-4.36 0"/>`), g.Group(children),
	)
}

func ProgrammingModuleCubeCodeModuleProgrammingPlugin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.39 3a.47.47 0 0 0-.21-.16l-6-2.27a.45.45 0 0 0-.36 0l-6 2.31A.47.47 0 0 0 .61 3a.48.48 0 0 0-.11.3v7.32a.5.5 0 0 0 .32.46l6 2.31h.36l6-2.31a.5.5 0 0 0 .32-.46V3.34a.48.48 0 0 0-.11-.34Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.46V5.5m0 0v7.96M.61 3.04L7 5.5l6.39-2.46"/>`), g.Group(children),
	)
}

func ProgrammingModulePuzzleOneCodePuzzleModuleProgrammingPluginPiece(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.09 1.19a2 2 0 0 1 .6 1.33L6.52.7a.67.67 0 0 1 1 0L9 2.24a2.18 2.18 0 0 0-.57.4a2.06 2.06 0 0 0 2.91 2.91a2.18 2.18 0 0 0 .4-.57l1.56 1.54a.67.67 0 0 1 0 1l-1.82 1.79a2.05 2.05 0 1 1-2.17 2.17L7.48 13.3a.67.67 0 0 1-1 0L5 11.76a2.18 2.18 0 0 0 .57-.4a2.06 2.06 0 0 0-2.93-2.91a2.18 2.18 0 0 0-.4.57L.7 7.48a.67.67 0 0 1 0-1l1.82-1.79a2 2 0 0 1-1.33-.6a2.05 2.05 0 0 1 2.9-2.9Z"/>`), g.Group(children),
	)
}

func ProgrammingModuleThreeCodeThreeModuleProgrammingPlugin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5.25" height="5.25" x=".5" y="8.25" rx=".5"/><rect width="5.25" height="5.25" x="8.25" y="8.25" rx=".5"/><rect width="5.25" height="5.25" x="4.38" y=".5" rx=".5"/></g>`), g.Group(children),
	)
}

func ProgrammingRssSquareWirelessRssFeedSquareTransmitBroadcast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="4" cy="10" r=".5"/><path d="M4.5 6.5a3 3 0 0 1 3 3m-3-6a6 6 0 0 1 6 6"/></g>`), g.Group(children),
	)
}

func ProgrammingRssSymbolWirelessFeedRssTransmitBroadcast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.13" cy="11.88" r="1.5"/><path d="M13.38 12.12A11.5 11.5 0 0 0 1.88.62m.24 4.76a6.5 6.5 0 0 1 6.5 6.5"/></g>`), g.Group(children),
	)
}

func ProgrammingScriptCodeCodeAngleProgrammingFileBracket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 10.5L.5 7L4 3.5m6 7L13.5 7L10 3.5"/>`), g.Group(children),
	)
}

func ProgrammingScriptFileCodeOneCodeFilesAngleProgrammingFileBracket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m2.5 13.5l-2-2l2-2m3 4l2-2l-2-2"/><path d="M2.5 7.5v-6a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1h-3"/><path d="M8.5.5v5h5"/></g>`), g.Group(children),
	)
}

func ProgrammingScriptFileCodeTwoCodeProgrammingTerminalShellFileLineCommandFiles(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 13.5l2-2l-2-2m2-2v-6a1 1 0 0 1 1-1h5l5 5v7a1 1 0 0 1-1 1H10"/><path d="M8.5.5v5h5m-9 8h3"/></g>`), g.Group(children),
	)
}

func ProgrammingScriptHtmlFiveLanguageFiveCodeProgrammingHtml(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12 10l-5 3.5L2 10L.5.5h13L12 10z"/><path d="M9.5 3.5h-5L5 6h4v2.5L7 10L5 8.5"/></g>`), g.Group(children),
	)
}

func ProgrammingScriptOneLanguageProgrammingCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 10.5v-8a2 2 0 0 1 2-2h6.25M10 3.5v8a2 2 0 0 1-2 2m-2-2V11a.5.5 0 0 0-.5-.5H1a.5.5 0 0 0-.5.5v.5a2 2 0 0 0 2 2H8a2 2 0 0 1-2-2Z"/><path d="M11.75.5h0a1.75 1.75 0 0 1 1.75 1.75V3a.5.5 0 0 1-.5.5h-3h0V2.25A1.75 1.75 0 0 1 11.75.5ZM6 4h1.5M6 7h1.5"/></g>`), g.Group(children),
	)
}

func ProgrammingScriptTwoLanguageProgrammingCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 3.5v8a2 2 0 0 1-2 2H.5a2 2 0 0 0 2-2v-9a2 2 0 0 1 2-2h7.25"/><path d="M11.75.5h0a1.75 1.75 0 0 1 1.75 1.75V3a.5.5 0 0 1-.5.5h-3h0V2.25A1.75 1.75 0 0 1 11.75.5ZM6 4h1.5M5 7h2.5M5 10h2.5"/></g>`), g.Group(children),
	)
}

func ProgrammingWebServerWorldInternetEarthWwwGlobeWorldwideWebNetwork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M.5 7h13m-4 0A11.22 11.22 0 0 1 7 13.5A11.22 11.22 0 0 1 4.5 7A11.22 11.22 0 0 1 7 .5A11.22 11.22 0 0 1 9.5 7Z"/></g>`), g.Group(children),
	)
}

func ReligionChristianityReligionJesusChristianityChristFishCulture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 4s-3.67 6-8 6a5.6 5.6 0 0 1-5-3a5.6 5.6 0 0 1 5-3c4.33 0 8 6 8 6"/>`), g.Group(children),
	)
}

func ReligionCrossOneReligionCrossCultureBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 3.5H9v-3H5v3H2v4h3v6h4v-6h3v-4z"/>`), g.Group(children),
	)
}

func ReligionCrossTwoReligionCrossCultureThinLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v13m-4-9h8"/>`), g.Group(children),
	)
}

func ReligionHexagramStarJewJewishJudaismHexagramCultureReligionDavid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 3.5h13L7 13.5L.5 3.5z"/><path d="M.5 10.5h13L7 .5l-6.5 10z"/></g>`), g.Group(children),
	)
}

func ReligionIslamReligionIslamMoonCrescentMuslimCultureStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.21 7a5.5 5.5 0 0 1 5.47-5.5a6.5 6.5 0 1 0 0 11A5.5 5.5 0 0 1 5.21 7Z"/><path d="m10.57 4.04l.91 1.81h1.81l-1.36 1.4l.43 2l-1.79-1l-1.71 1l.36-2l-1.36-1.4h1.81l.9-1.81z"/></g>`), g.Group(children),
	)
}

func ReligionShintoReligionGateCultureShintoJapanJapaneseShrine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 4a.5.5 0 0 1-.5.5H1A.5.5 0 0 1 .5 4V.5a11.55 11.55 0 0 0 13 0ZM3 4.5v9m8-9v9M.5 8h13M7 4.5V8"/>`), g.Group(children),
	)
}

func ReligionSymbolPeaceReligionPeaceWarCulture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5v13m-5.07-2.43L7 6l5.07 5.07"/></g>`), g.Group(children),
	)
}

func ReligionSymbolYinYangReligionTaoYinYangTaoismCulture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5S2.93 4.09 7 7c3.5 2.5 0 6.5 0 6.5"/><circle cx="9" cy="4.5" r="1"/><circle cx="5" cy="9.5" r="1"/></g>`), g.Group(children),
	)
}

func ShippingBoxFragileFragileShippingGlassDeliveryWineCrackShipmentSignSticker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.67.5a1 1 0 0 0-1 .89C2.56 2.45 2.5 4.05 2.5 4.5c0 2.62 2 4 4.5 4s4.5-1.38 4.5-4c0-.45-.06-2-.17-3.11a1 1 0 0 0-1-.89ZM7 8.5v5m-2 0h4"/><path d="m7 .5l-1.5 2l2 1.5l-1 1.5"/></g>`), g.Group(children),
	)
}

func ShippingBoxOneBoxPackageLabelDeliveryShipmentShipping(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M9 .5v5a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-5M8.5 11H11"/></g>`), g.Group(children),
	)
}

func ShippingBoxTwoBoxPackageLabelDeliveryShipmentShippingThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 .5v4M8.5 11H11M.5 4.5h13v8a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-8h0Z"/><path d="M.5 4.5L2 1.61A2 2 0 0 1 3.74.5h6.52a2 2 0 0 1 1.79 1.11L13.5 4.5"/></g>`), g.Group(children),
	)
}

func ShippingTransferCartPackageBoxFulfillmentCartWarehouseShippingDelivery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6.5" height="6.5" x="5.5" y="3" rx="1"/><path d="M4 12h7.61a1 1 0 0 1 .7.29l1.19 1.21M.5.5h1a1 1 0 0 1 1 1V10M4 11.75A2.11 2.11 0 0 1 4 12a1.74 1.74 0 1 1 0-.25ZM8.5 7h1"/></g>`), g.Group(children),
	)
}

func ShippingTransferTruckTimeTruckShippingDeliveryTimeWaitingDelay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6" cy="12.49" r="1"/><circle cx="10.5" cy="12.49" r="1"/><path d="M8.38 3.53A4 4 0 1 0 2 7.62m2.5-3.11L6 3.01m.5 6.49v-1H5a2 2 0 0 0-2 2v2"/><path d="M13.5 12.49v-5a1 1 0 0 0-1-1h-5a1 1 0 0 0-1 1v2"/></g>`), g.Group(children),
	)
}

func ShippingTransferTruckTruckShippingDelivery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 5.5h-3a2 2 0 0 0-2 2v4H2m10.5 0h1v-7a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v5.68M6 11.5h2.5"/><circle cx="4" cy="11.5" r="2"/><circle cx="10.5" cy="11.5" r="2"/></g>`), g.Group(children),
	)
}

func ShippingWarehouseDeliveryWarehouseShippingFulfillment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 4.5h13V13a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5V4.5h0Zm0 0h0a4 4 0 0 1 4-4h5a4 4 0 0 1 4 4h0"/><path d="M11 13.5V8a.5.5 0 0 0-.5-.5h-7A.5.5 0 0 0 3 8v5.5m0-4h8m-8 2h8"/></g>`), g.Group(children),
	)
}

func ShoppingBagHandBagOneShoppingBagPurseGoodsItemProducts(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.27 13.5H2.73a2 2 0 0 1-2-2.22l.67-5.89a1 1 0 0 1 1-.89h9.2a1 1 0 0 1 1 .89l.65 5.89a2 2 0 0 1-1.98 2.22Z"/><path d="M3 4.5a4 4 0 0 1 8 0m-6.5 3h5"/></g>`), g.Group(children),
	)
}

func ShoppingBagHandBagTwoShoppingBagPurseGoodsItemProducts(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.88 12.39a1 1 0 0 1-.25.78a1 1 0 0 1-.75.33H2.12a1 1 0 0 1-.75-.33a1 1 0 0 1-.25-.78L2 4.5h10ZM4.5 4.5V3a2.5 2.5 0 0 1 5 0v1.5"/>`), g.Group(children),
	)
}

func ShoppingBagSuitcaseOneProductBusinessBriefcase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10" x=".5" y="3.5" rx="1"/><path d="M5 .5h4a1 1 0 0 1 1 1v2h0h-6h0v-2a1 1 0 0 1 1-1ZM3.5 7h7m-7 3h7"/></g>`), g.Group(children),
	)
}

func ShoppingBuildingRealHomeTowerBuildingHouseEstate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 13.5h-8V4l4-3.5l4 3.5v9.5zm0 0h5v-7h-5m-4 7v-2M3 8.5h3m-3-3h3"/>`), g.Group(children),
	)
}

func ShoppingCartBasketOneShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.36 6H1.64a1 1 0 0 0-1 1.13l.73 5.5a1 1 0 0 0 1 .87h9.24a1 1 0 0 0 1-.87l.73-5.5A1 1 0 0 0 12.36 6ZM4.5 8.5V11M7 8.5V11m2.5-2.5V11"/><path d="M9.48 1.54A2.79 2.79 0 0 1 11.78 4L12 6M2 6l.22-2a2.79 2.79 0 0 1 2.3-2.44"/><path d="M9.5 1.75A1.25 1.25 0 0 1 8.25 3h-2.5a1.25 1.25 0 0 1 0-2.5h2.5A1.25 1.25 0 0 1 9.5 1.75Z"/></g>`), g.Group(children),
	)
}

func ShoppingCartBasketThreeShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.36 6.62a1 1 0 0 0-.24-.78a1 1 0 0 0-.75-.34H1.63a1 1 0 0 0-.75.34a1 1 0 0 0-.24.78l.75 6a1 1 0 0 0 1 .88h9.24a1 1 0 0 0 1-.88ZM2.5 5.5V5a4.5 4.5 0 0 1 9 0v.5M5 8.5v2m4-2v2"/>`), g.Group(children),
	)
}

func ShoppingCartBasketTwoShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.28 6H1.72a1 1 0 0 0-1 1.2l1.1 5.5a1 1 0 0 0 1 .8h8.36a1 1 0 0 0 1-.8l1.1-5.5a1 1 0 0 0-1-1.2ZM9 2.5L11 6M3 6l2-3.5"/>`), g.Group(children),
	)
}

func ShoppingCartOneShoppingCartCheckout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5H11l-.87 8.65a1 1 0 0 1-1 .85h-6.3a1 1 0 0 1-1-.68l-1.33-4a1 1 0 0 1 .14-.9A1 1 0 0 1 1.5 4h9.15"/><circle cx="3" cy="13" r=".5"/><circle cx="9.5" cy="13" r=".5"/></g>`), g.Group(children),
	)
}

func ShoppingCartThreeShoppingCartCheckout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.75 4.5l.25 1m3.25-1L7 5.5m-5.18 2H10l1-5H1a.5.5 0 0 0-.49.59l.82 4a.49.49 0 0 0 .49.41Zm9.18-5l.42-1.6a.5.5 0 0 1 .49-.4h1.59m-3.5 7l-.42 2.1a.5.5 0 0 1-.49.4H3"/><circle cx="3.5" cy="13" r=".5"/><circle cx="8.5" cy="13" r=".5"/></g>`), g.Group(children),
	)
}

func ShoppingCartTwoShoppingCartCheckout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.82 7.5H10l1-5H1a.5.5 0 0 0-.49.59l.82 4a.49.49 0 0 0 .49.41Zm9.18-5l.42-1.6a.5.5 0 0 1 .49-.4h1.59m-3.5 7l-.42 2.1a.5.5 0 0 1-.49.4H3"/><circle cx="3.5" cy="13" r=".5"/><circle cx="8.5" cy="13" r=".5"/></g>`), g.Group(children),
	)
}

func ShoppingCatergoriesBabyBotlleBottleMilkFamilyChildrenFormulaCareChildKidBaby(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 6.5v6a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-6m0 0v-2a1 1 0 0 0-1-1a1 1 0 0 1-1-1V2a1.5 1.5 0 0 0-3 0v.5a1 1 0 0 1-1 1a1 1 0 0 0-1 1v2m-1 0h9m-4 3h3"/>`), g.Group(children),
	)
}

func ShoppingCatergoriesBallSportsBallSportBasketball(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5v13m-4.9-2.23A5 5 0 0 0 4.5 7a5 5 0 0 0-2.4-4.27m9.8 0a5 5 0 0 0 0 8.54"/></g>`), g.Group(children),
	)
}

func ShoppingCatergoriesChairDesignLoungeFurnitureChairInteriorDecorateArmchairDecoration(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 7.5h1a1 1 0 0 1 1 1v5h0h-3h0v-5a1 1 0 0 1 1-1Zm10 0h1a1 1 0 0 1 1 1v5h0h-3h0v-5a1 1 0 0 1 1-1Zm-8 6h7m-8-6v-2a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3v2m-8 2h7"/>`), g.Group(children),
	)
}

func ShoppingCatergoriesDressClothingDressSkirtWomen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 3.5L11 5l2-2L10.5.5h-7L1 3l2 2l1.5-1.5V7c0 1.35-1.62 2.3-1.94 5.4a1 1 0 0 0 .25.77a1 1 0 0 0 .74.33h6.9a1 1 0 0 0 .74-.33a1 1 0 0 0 .25-.77C11.12 9.3 9.5 8.35 9.5 7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.25a6.47 6.47 0 0 0 5 0"/>`), g.Group(children),
	)
}

func ShoppingCatergoriesPhoneAndroidPhoneMobileDeviceSmartphoneIphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="13" x="2" y=".5" rx="1"/><path d="M4.5 3h5v5h-5z"/><circle cx="7" cy="10.75" r=".5"/></g>`), g.Group(children),
	)
}

func ShoppingCatergoriesRingMoneyDiamondPaymentWealthFinanceRingAccessories(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="9" r="4.5"/><path d="m8.5.5l1 1l-2.5 3l-2.5-3l1-1h3z"/></g>`), g.Group(children),
	)
}

func ShoppingCatergoriesShirtClothingTShirtMenTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 1.5l3 3l-2 2l-1-1v6a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1v-6l-1 1l-2-2l3-3Z"/>`), g.Group(children),
	)
}

func ShoppingCatergoriesSofaDecorationFurnitureInteriorDesignCouchSofaDecorate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="7" x="3" y="1.5" rx="1"/><path d="M3 4.5H1.5a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1H11M4 11l-.5 2.5M10 11l.5 2.5"/></g>`), g.Group(children),
	)
}

func ShoppingGiftRewardBoxSocialPresentGiftMediaRatingBow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="5" x=".5" y="3" rx="1"/><path d="M12.5 8v4.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1V8M7 3v10.5m3-13L7 3L4 .5"/></g>`), g.Group(children),
	)
}

func ShoppingJewelryDiamondOneDiamondMoneyPaymentFinanceWealth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 5.5l-6.5 7l-6.5-7l3-4h7l3 4zm-13 0h13"/><path d="m3.5 1.5l3.5 4l3.5-4M7 5.5v7"/></g>`), g.Group(children),
	)
}

func ShoppingJewelryDiamondTwoDiamondMoneyPaymentFinanceWealth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.64 1.54H3.36a1.07 1.07 0 0 0-.85.46L.69 4.52a1.05 1.05 0 0 0 .06 1.29l5.46 6.29a1 1 0 0 0 1.58 0l5.46-6.29a1.05 1.05 0 0 0 .06-1.29L11.49 2a1.07 1.07 0 0 0-.85-.46Z"/><path d="M6.48 1.53L4.04 5.31L7 12.46m.55-10.93l2.43 3.78L7 12.46M.52 5.31h12.96"/></g>`), g.Group(children),
	)
}

func ShoppingJewelryGoldGoldMoneyPaymentBarsFinanceWealthBullion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.66 9.3a1 1 0 0 0-1-.8H2.32a1 1 0 0 0-1 .8L.5 13.5h6Zm7 0a1 1 0 0 0-1-.8H9.32a1 1 0 0 0-1 .8l-.82 4.2h6Z"/><path d="m10 8.5l-.84-4.2a1 1 0 0 0-1-.8H5.82a1 1 0 0 0-1 .8L4 8.5"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.68 8.5h4.64"/>`), g.Group(children),
	)
}

func ShoppingStoreOneStoreShopShopsStores(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 8.5V13a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8.5M8 8.5v5M1.5 10H8M.5 4L2 .5h10L13.5 4H.5zm4.25 0v1a2 2 0 0 1-2 2h-.28a2 2 0 0 1-2-2V4m8.78 0v1a2 2 0 0 1-2 2h-.5a2 2 0 0 1-2-2V4m8.75 0v1a2 2 0 0 1-2 2h-.25a2 2 0 0 1-2-2V4"/>`), g.Group(children),
	)
}

func ShoppingStoreSignageOneSignStoreShopShopsSignageStores(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="2"/><path d="m4 5.5l3-5l3 5"/></g>`), g.Group(children),
	)
}

func ShoppingStoreSignageThreeStreetSandwichShopsShopStoresBoardSignStore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 2.5h-7l-3 11h7l3-11zm3 11l-3-11m-2.18 8h4.36"/>`), g.Group(children),
	)
}

func ShoppingStoreSignageTwoSignStoreShopShopsSignageStoresPole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y=".5" rx="2"/><path d="M7 13.5v-6"/></g>`), g.Group(children),
	)
}

func ShoppingStoreTwoStoreShopShopsStores(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 8.5V13a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8.5M13 6H1a.5.5 0 0 1-.5-.5v-2l1.22-2.45a1 1 0 0 1 .9-.55h8.76a1 1 0 0 1 .9.55L13.5 3.5v2a.5.5 0 0 1-.5.5ZM8 8.5v5M1.5 10H8M.5 3.5h13"/>`), g.Group(children),
	)
}

func TravelAirportArrivalTimePlaneAirplaneTripLandTravelTimeAdventureTimerClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.18 6.08A4 4 0 1 0 2.52 8M4.5 4.5L6 3m6.3 4.09l-1.06.29a.34.34 0 0 0-.24.32V9l-4.4.52a2.48 2.48 0 0 0-2.31 3a.66.66 0 0 0 .83.47l2.65-.71l.23-.02l1.33 1.15a.39.39 0 0 0 .34.08l1.2-.33a.36.36 0 0 0 .17-.61l-.83-.88l.17-.05l2.59-.7a.67.67 0 0 0 .49-.8l-.76-2.81a.34.34 0 0 0-.4-.22Z"/>`), g.Group(children),
	)
}

func TravelAirportBaggageCheckBaggageTravelAdventureLuggageBagChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9.5" x=".5" y="4" rx="2"/><path d="M4 13.5V4m6 9.5V4M4.5 4a2.5 2.5 0 0 1 5 0"/></g>`), g.Group(children),
	)
}

func TravelAirportDepartureTimeTravelPlaneTripAirplaneTimeOffAdventureTimerTakeClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.82 6.08A4 4 0 1 1 13.5 4.5a4 4 0 0 1-1.17 2.83M9.5 4.5L11 3m-.74 6.16l-1-.34a.34.34 0 0 0-.44.18l-.73 1.13l-4-2A2.49 2.49 0 0 0 .53 9.39a.68.68 0 0 0 .45.85l2.61.84l.26.08l.49 1.69a.36.36 0 0 0 .24.25l1.18.38a.37.37 0 0 0 .48-.42L6 11.87l.17.05l2.55.83a.67.67 0 0 0 .85-.41l.9-2.77a.34.34 0 0 0-.21-.41Z"/>`), g.Group(children),
	)
}

func TravelAirportEarthAirplaneTravelPlaneTripAirplaneInternationalAdventureGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 8.5a5 5 0 1 1-8.25-3.8"/><path d="M5.28 6.65a1.6 1.6 0 0 0-.16.7a1.52 1.52 0 0 0 1.53 1.53a.77.77 0 0 1 .77.77v3.47M.57 9.27h1.85A1.54 1.54 0 0 1 4 10.81v2.45m9.26-11.39l-1-.34a.34.34 0 0 0-.39.13l-.73 1.13l-4-2A2.49 2.49 0 0 0 3.53 2.1A.68.68 0 0 0 4 3l2.61.84l.26.09l.49 1.68a.36.36 0 0 0 .24.25l1.18.38a.37.37 0 0 0 .48-.41L9 4.58h.17l2.55.83a.67.67 0 0 0 .85-.41l.9-2.77a.34.34 0 0 0-.21-.36Z"/></g>`), g.Group(children),
	)
}

func TravelAirportLandingLandPlaneTravelAdventureAirplane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 1.93l1.38.54a.46.46 0 0 1 .3.47l-.18 1.8l5.84 1.35A3.37 3.37 0 0 1 13 10.5a.91.91 0 0 1-1.19.5L8.36 9.65L8 9.52l-2 1.34a.52.52 0 0 1-.47.05L4 10.3a.5.5 0 0 1-.14-.85l1.28-1.06l-.23-.09L1.54 7A.9.9 0 0 1 1 5.83l1.44-3.66A.44.44 0 0 1 3 1.93ZM.5 13.5h13"/>`), g.Group(children),
	)
}

func TravelAirportPassportTravelBookIdAdventureVisa(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><circle cx="7" cy="6" r="3"/><path d="M4 6h6M7 9V3"/></g>`), g.Group(children),
	)
}

func TravelAirportTakeOffTravelPlaneAdventureAirplaneTakeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.19 3.78l-1.37-.44a.46.46 0 0 0-.51.17L10.37 5L5.21 2.43A3.24 3.24 0 0 0 .54 4.08a.88.88 0 0 0 .58 1.1l3.39 1.1l.34.11l.64 2.19a.45.45 0 0 0 .31.32l1.54.5A.48.48 0 0 0 8 8.86L7.68 7.3l.22.07l3.31 1.07a.86.86 0 0 0 1.11-.52l1.16-3.6a.44.44 0 0 0-.29-.54ZM.5 13.5h13"/>`), g.Group(children),
	)
}

func TravelHotelAirConditionerHeatingAcAirHvacCoolCoolingColdHotConditioning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y=".5" rx="1"/><path d="M11 7.5v-2a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v2m0 3l-.5 3m4.5-3v3m4-3l.5 3"/></g>`), g.Group(children),
	)
}

func TravelHotelBathTubBatheBathBathroomTowelWaterTub(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 4.5h-5a1 1 0 0 0-1 1v1a4 4 0 0 0 4 4h5a4 4 0 0 0 4-4v-1a1 1 0 0 0-1-1h-2"/><path d="M6.5 3.5h4v4h-4z"/></g>`), g.Group(children),
	)
}

func TravelHotelBedOneBedSingleBedroomsBedroomHotel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="M9.5.5v3h-5v-3M1.5 6h11"/></g>`), g.Group(children),
	)
}

func TravelHotelBedTwoBedDoubleBedroomBedroomsQueenKingFullHotel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M7 3.5H3v-3m8 0v3H7v-3M.5 6h13"/></g>`), g.Group(children),
	)
}

func TravelHotelBellServiceConciergePorterCallRingBellhopBellReception(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 10.25a6.5 6.5 0 0 1 13 0Zm0 2.5h13m-6.5-9v-2.5m-1.5 0h3"/>`), g.Group(children),
	)
}

func TravelHotelDumbellSportsWeightsDumbbellSportFitness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4" height="6" x=".5" y="4" rx="1"/><rect width="4" height="6" x="9.5" y="4" rx="1"/><path d="M4.5 7h5"/></g>`), g.Group(children),
	)
}

func TravelHotelHangerHangerLockerCheckCoatRoomCloak(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 3a2 2 0 1 0-2 2v1m0 0l-6 4.6a1.32 1.32 0 0 0-.5 1.06A1.34 1.34 0 0 0 1.84 13h10.32a1.34 1.34 0 0 0 1.34-1.34a1.32 1.32 0 0 0-.5-1.06Z"/>`), g.Group(children),
	)
}

func TravelHotelIronLaundryIronHeat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.59 5.5H8.91a8.78 8.78 0 0 0-8.32 6h11.82a1 1 0 0 0 .76-.35a1 1 0 0 0 .23-.81Zm0 0l-.22-1.33a2 2 0 0 0-2-1.67H4.59"/>`), g.Group(children),
	)
}

func TravelHotelLaundryBasketBathroomLaundryBinBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 1v9a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3V1M.5 1h13m-12 3h3m-3 3h3m-3 3h3m5-6h3m-3 3h3m-3 3h3"/>`), g.Group(children),
	)
}

func TravelHotelLaundryLaundryMachine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="7" cy="8" r="3.5"/><path d="M.5 3.5h3m7 0h3M7 7.97v3.5M4 6.2l3 1.77m3-1.77L7 7.97"/></g>`), g.Group(children),
	)
}

func TravelHotelOneStarOneStarReviewsReviewRatingHotelStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6 1.05a1.18 1.18 0 0 1 2 0l1.94 3.07l2.53.31a1.16 1.16 0 0 1 .75 1.91l-2.11 2.42l.68 3.35a1.17 1.17 0 0 1-.46 1.17a1.19 1.19 0 0 1-1.26.07L7 11.67l-3.07 1.68a1.19 1.19 0 0 1-1.26-.07a1.17 1.17 0 0 1-.46-1.17l.68-3.35L.78 6.34a1.16 1.16 0 0 1 .75-1.91l2.53-.31Z"/><path d="M7 9.52v-4H6m-.5 4h3"/></g>`), g.Group(children),
	)
}

func TravelHotelParkingSignDiscountCouponParkingPricePrices(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="M5.5 3.5v7m0-3h2a2 2 0 0 0 0-4h-2"/></g>`), g.Group(children),
	)
}

func TravelHotelPetPawPawFootAnimalsPetsFootprintTrack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="1.5" cy="8" rx="1" ry="1.5"/><ellipse cx="4.5" cy="3.5" rx="1" ry="1.5"/><ellipse cx="9.5" cy="3.5" rx="1" ry="1.5"/><ellipse cx="12.5" cy="8" rx="1" ry="1.5"/><path d="M10 10c0 1.38-1.62 2-3 2s-3-.62-3-2s1-3.5 3-3.5s3 2.12 3 3.5Z"/></g>`), g.Group(children),
	)
}

func TravelHotelPoolLadderTwoPoolStairsSwimSwimmingWaterLadder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5H13a2 2 0 0 1-2-2a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-2 2H.5m2-11a2 2 0 0 1 4 0v7m3-9a2 2 0 0 1 2 2v7m-5-5h5m-5 3h5"/>`), g.Group(children),
	)
}

func TravelHotelServingDomeCookToolDomeKitchenDrinkServingPlatterDishToolsFoodCooking(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 9.75v-.5a5.5 5.5 0 0 0-11 0v.5m12 0H.5l.32 1.07a2 2 0 0 0 1.92 1.43h8.52a2 2 0 0 0 1.92-1.43Zm-6.5-6v-2"/>`), g.Group(children),
	)
}

func TravelHotelServingDomeHandPorterServiceRoomPlateHandBellhopPlatterGiveFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 5.5a5 5 0 0 1 10 0m-11 0h12m-13 3H4A2.5 2.5 0 0 1 6.5 11h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5H.5"/>`), g.Group(children),
	)
}

func TravelHotelShowerHeadBatheBathBathroomShowerWaterHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 6.5a4 4 0 0 1 8 0Zm2 3v1m-1.5 2v1m3.5-1v1m3.5-1v1M9 9.5v1m-2-8v-2"/>`), g.Group(children),
	)
}

func TravelHotelSpaSpaFlowerLotusLily(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 10.5c0 1.66-2.91 3-6.5 3s-6.5-1.34-6.5-3m9-5C9.5 6.88 7.5 9 7 9S4.5 6.88 4.5 5.5c0-2 2.5-5 2.5-5s2.5 3 2.5 5Z"/><path d="M4.88 3.89A15.7 15.7 0 0 0 1 3s.35 3.89 1.77 5.3c1 1 3.89 1.06 4.24.71"/><path d="M9.12 3.89A15.7 15.7 0 0 1 13 3s-.35 3.89-1.77 5.3c-1 1-3.89 1.06-4.24.71"/></g>`), g.Group(children),
	)
}

func TravelHotelThreeStarThreeStarsReviewsReviewRatingHotelStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6 1.05a1.18 1.18 0 0 1 2 0l1.94 3.07l2.53.31a1.16 1.16 0 0 1 .75 1.91l-2.11 2.42l.68 3.35a1.17 1.17 0 0 1-.46 1.17a1.19 1.19 0 0 1-1.26.07L7 11.67l-3.07 1.68a1.19 1.19 0 0 1-1.26-.07a1.17 1.17 0 0 1-.46-1.17l.68-3.35L.78 6.34a1.16 1.16 0 0 1 .75-1.91l2.53-.31Z"/><path d="M5.5 5.52h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2m3-2h-3"/></g>`), g.Group(children),
	)
}

func TravelHotelTowelTowelSpaWetWaterHot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 10a3 3 0 0 0-3-3h-9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h9a3 3 0 0 0 3-3Zm-13 0h9m-6-3a3 3 0 0 1 0-6h9a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-1.86m2.86-3h-9"/>`), g.Group(children),
	)
}

func TravelHotelTwoStarTwoStarsReviewsReviewRatingHotelStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6 1.05a1.18 1.18 0 0 1 2 0l1.94 3.07l2.53.31a1.16 1.16 0 0 1 .75 1.91l-2.11 2.42l.68 3.35a1.17 1.17 0 0 1-.46 1.17a1.19 1.19 0 0 1-1.26.07L7 11.67l-3.07 1.68a1.19 1.19 0 0 1-1.26-.07a1.17 1.17 0 0 1-.46-1.17l.68-3.35L.78 6.34a1.16 1.16 0 0 1 .75-1.91l2.53-.31Z"/><path d="M5.5 5.52h2a1 1 0 0 1 0 2h-1a1 1 0 0 0-1 1v1h3"/></g>`), g.Group(children),
	)
}

func TravelMapEarthOnePlanetEarthGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M1 9.5h1.75A1.75 1.75 0 0 0 4.5 7.75v-1.5A1.75 1.75 0 0 1 6.25 4.5A1.75 1.75 0 0 0 8 2.75V.57m5.5 6.33a3.56 3.56 0 0 0-1.62-.4H9.75a1.75 1.75 0 0 0 0 3.5A1.25 1.25 0 0 1 11 11.25v.87"/></g>`), g.Group(children),
	)
}

func TravelMapEarthTwoPlanetEarthGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M12.48 3.5h-4a2 2 0 0 0 0 4a1 1 0 0 1 1 1V13M.58 8H3a2 2 0 0 1 2 2v3.19"/></g>`), g.Group(children),
	)
}

func TravelMapFlagNavigationMapMapsFlagGpsLocationDestinationGoal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v13m9-6h-9v-7h9L8.5 4l3 3.5z"/>`), g.Group(children),
	)
}

func TravelMapGlobeModelPlanetEarthGlobeWorld(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.16" cy="5.75" r="3"/><path d="M6.16 11v2.5m-1.5 0h3M6.16.5a5.25 5.25 0 1 1-3.57 9.1"/></g>`), g.Group(children),
	)
}

func TravelMapLocationCompassOneArrowCompassLocationGpsMapMapsPoint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m7.5 10.5l2-6l-6 2L6 8l1.5 2.5z"/></g>`), g.Group(children),
	)
}

func TravelMapLocationPinNavigationMapMapsPinGpsLocation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 5c0 2.49-4.5 8.5-4.5 8.5S2.5 7.49 2.5 5a4.5 4.5 0 0 1 9 0Z"/><circle cx="7" cy="5" r="1.5"/></g>`), g.Group(children),
	)
}

func TravelMapLocationTargetOffTargetLocationMapOffServicesNavigationMapsGps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 2.4V.5m0 13v-2M11.6 7h1.9m-12-5.5l11 11M.5 7h2m1.3-3.2c1.7-1.8 4.6-1.9 6.4-.2s1.9 4.5.2 6.4c-.1.1-.1.2-.2.2M8 11.4c-2.4.5-4.8-1-5.4-3.4c-.1-.3-.1-.7-.1-1c0-.5.1-1 .2-1.4"/>`), g.Group(children),
	)
}

func TravelMapLocationTargetOneNavigationLocationMapServicesMapsGpsTarget(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="4.5"/><circle cx="7" cy="7" r=".5"/><path d="M7 2.5v-2m0 13v-2M11.5 7h2M.5 7h2"/></g>`), g.Group(children),
	)
}

func TravelMapLocationTargetTwoNavigationLocationMapServicesMapsGpsTarget(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="4.5"/><path d="M7 2.5v-2m0 13v-2M11.5 7h2M.5 7h2"/></g>`), g.Group(children),
	)
}

func TravelMapNavigationArrowCompassArrowMapBearingNavigationMapsHeadingGps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 13.5l5-13l-13 5l6 2l2 6z"/>`), g.Group(children),
	)
}

func TravelMapNavigationArrowOffCompassArrowMapBearingNavigationMapsHeadingGpsOffDisable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.28 3.28L.5 5.5l6 2l2 6l2.22-5.78m1.08-2.8L13.5.5L8.91 2.27M3.5.5l10 10"/>`), g.Group(children),
	)
}

func TravelMapNavigationMapMapsGps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.83 12.5l-4.33 1v-12l4.33-1v12zm0 0l4.34 1v-12L4.83.5v12zm8.67 0l-4.33 1v-12l4.33-1v12z"/>`), g.Group(children),
	)
}

func TravelMapRectangleFlagNavigationMapMapsFlagGpsLocationDestinationGoal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v13m0-13h9v7h-9"/>`), g.Group(children),
	)
}

func TravelMapTriangleFlagNavigationMapMapsFlagGpsLocationDestinationGoal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v13m0-13l9 4.5l-9 4.5"/>`), g.Group(children),
	)
}

func TravelPlacesAnchorAnchorMarinaHarborPort(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 8H1.5a5.5 5.5 0 0 0 11 0H11"/><circle cx="7" cy="2" r="1.5"/><path d="M7 3.5v10m-1.5-7h3"/></g>`), g.Group(children),
	)
}

func TravelPlacesBeachIslandWavesOutdoorRecreationTreeBeachPalmWaveWater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 13.48H13a2 2 0 0 1-2-2a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-2 2H.5m9.5-4a5.49 5.49 0 0 0-8.48 0"/><path d="M6.5 7.53c.06-2.26.75-4.32 2.25-5.06M5.76.57a2.58 2.58 0 0 1 3 1.9"/><path d="M12.41 2.84a2.78 2.78 0 0 0-3.66-.37"/><path d="M5.08 3.54a3 3 0 0 1 3.67-1.07a2.55 2.55 0 0 1 1.89 3"/></g>`), g.Group(children),
	)
}

func TravelPlacesCampingTentOutdoorRecreationCampingTentTeepeeTipi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 10.5v3m2-13l-8.5 13h13L5 .5"/>`), g.Group(children),
	)
}

func TravelPlacesCartCartTravelLocalBuckboardRetro(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.59 10H4V3a1 1 0 0 1 1-1h4.5a4 4 0 0 1 4 4v3a1 1 0 0 1-1 1h-2.09M4 10H.5V8"/><circle cx="9" cy="10.5" r="1.5"/><path d="M7.5 5h3"/></g>`), g.Group(children),
	)
}

func TravelPlacesColumnOnePillarBuildingHistoricalColumn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 13.25v-7.5m2.5 7.5v-5.5m3 5.5v-5.5m2.5 5.5v-7.5m0-5H3a2.5 2.5 0 1 0 2.29 3.5h3.42A2.5 2.5 0 1 0 11 .75Z"/>`), g.Group(children),
	)
}

func TravelPlacesColumnTwoPillarBuildingHistoricalColumn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.75" cy="2.75" r="2.25"/><circle cx="11.25" cy="2.75" r="2.25"/><path d="M3 .5h8m-6.84 4h5.68M3 13.5V4.99m2.5 8.51v-6m3 6v-6m2.5 6V4.99"/></g>`), g.Group(children),
	)
}

func TravelPlacesDiamondDiamondMoneyPaymentFinanceWealth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.89.93a1 1 0 0 0-.82-.43H3.93a1 1 0 0 0-.82.43l-1.6 2.31a1 1 0 0 0 .15 1.28l4.67 4.72a1 1 0 0 0 1.34 0l4.67-4.72a1 1 0 0 0 .15-1.28Z"/><path d="M7 .5L4.68 3.94l2.18 5.55M7 .5l2.32 3.44l-2.18 5.55M1.35 3.94h11.3M13 13.5h0a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2h0"/></g>`), g.Group(children),
	)
}

func TravelPlacesHotSpringRelaxLocationOutdoorRecreationSpa(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 7a2.8 2.8 0 0 1 1.5 2.24c0 1.93-2.91 3.5-6.5 3.5S.5 11.18.5 9.25A2.8 2.8 0 0 1 2 7"/><path d="M4.08 2.25c-2 2.5 2 3.5 0 6m2.92-7c-2 2.5 2 5.5 0 8m2.92-7c-2 2.5 2 3.5 0 6"/></g>`), g.Group(children),
	)
}

func TravelPlacesMountainSunLandmarkMountainPlace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 9.26l-3.79-5a1 1 0 0 0-1.42 0L.5 13.5h13"/><path d="m4.69 8.5l1.47 1a1 1 0 0 0 1.16 0l.89-.67a1 1 0 0 1 1.2 0l.89.67a1 1 0 0 0 1.15 0l1.48-1"/><circle cx="3" cy="3" r="2.5"/></g>`), g.Group(children),
	)
}

func TravelPlacesPaintingPaintingEntertainmentDisplayMuseumEventHobbyExhibit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="4.24" rx=".5"/><circle cx="4.25" cy="7.99" r="1.25"/><path d="m3.75 13.24l4.7-4a1.32 1.32 0 0 1 1.87.15l3.07 3.68M3.5 4.24L6.25 1.1a1 1 0 0 1 1.5 0l2.75 3.14"/></g>`), g.Group(children),
	)
}

func TravelPlacesSnorkleDivingScubaOutdoorRecreationOceanMaskWaterSeaSnorkle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 10.5a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V.5"/><path d="M10.5 5.5a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1l1.17-1.75a1 1 0 0 1 1.66 0L7.5 9.5h1a2 2 0 0 0 2-2Z"/></g>`), g.Group(children),
	)
}

func TravelPlacesStreetSignCrossroadStreetSignMetaphorDirections(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 4.5H2.5l-2-2l2-2H7m0 7h4.5l2-2l-2-2H7m0 7H2.5l-2-2l2-2H7m0-6v13"/>`), g.Group(children),
	)
}

func TravelPlacesTheaterMaskHobbyTheaterMasksDramaEventShowEntertainment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.77 8.5A5.91 5.91 0 0 1 .5 4.84V1.58a9.65 9.65 0 0 1 8.87 0"/><path d="M7.79 13.44h0a2.27 2.27 0 0 1-2-.51h0a6.66 6.66 0 0 1-2.11-6.67l.73-2.92a9.88 9.88 0 0 1 9.09 2.28l-.73 2.91a6.67 6.67 0 0 1-4.98 4.91Z"/><path d="M5.61 6.51a1 1 0 0 1 1-.27a1 1 0 0 1 .73.69m2.23.57a1 1 0 0 1 1.7.43"/></g>`), g.Group(children),
	)
}

func TravelPlacesTrafficSignStreetSignTrafficTravel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="4" r="3.5"/><path d="M7 7.5v6m-1.5 0h3"/></g>`), g.Group(children),
	)
}

func TravelTransportationAirplaneTravelPlaneAdventureAirplane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.75 2.75h-1.61a.49.49 0 0 0-.48.38l-.51 2l-5-1a3.69 3.69 0 0 0-4.4 3.59a1 1 0 0 0 1 1h4.4l1 1.58a2 2 0 0 0 1.68.92h1.1a.5.5 0 0 0 .44-.73l-.88-1.74h2.76a1 1 0 0 0 1-1v-4.5a.5.5 0 0 0-.5-.5Z"/>`), g.Group(children),
	)
}

func TravelTransportationBusTransportationTravelBusTransitTransportMotorcoachPublic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="11.5" x="1.5" y=".5" rx="1"/><path d="M3.5 12v1.5m7-1.5v1.5M1.5 7h11"/><circle cx="4" cy="9.5" r=".5"/><circle cx="10" cy="9.5" r=".5"/></g>`), g.Group(children),
	)
}

func TravelTransportationCarTransportationTravelTaxiTransportCab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 10.5a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H11l-.77-2.32a1 1 0 0 0-1-.68H5.22a1 1 0 0 0-1 .68L3.5 6.5h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1"/><circle cx="3.5" cy="10.5" r="2"/><circle cx="10.5" cy="10.5" r="2"/><path d="M5.5 10.5h3M7 3.5v-2"/></g>`), g.Group(children),
	)
}

func TravelTransportationSailShipTravelBoatTransportationTransportOceanShipSeaWater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 9.5H.5l1 2.7a2 2 0 0 0 1.88 1.3h7.22a2 2 0 0 0 1.88-1.3Zm-7 0v-9m0 0l-5 6h5m2-4l3.5 4H8.5"/>`), g.Group(children),
	)
}

func TravelWafinderSinkWashCleanToiletBathroomWater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5h13v2a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3v-2h0Zm7-6a2 2 0 0 0-4 0v6m4-4v1"/>`), g.Group(children),
	)
}

func TravelWayfinderDisabilityPersonAccessWheelchairAccomodationHumanDisabilityDisabledUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8" cy="2.5" r="2"/><path d="M8 4.5v4h3a1 1 0 0 1 1 1v4m-6.5-7a3.5 3.5 0 1 0 3.22 4.88"/></g>`), g.Group(children),
	)
}

func TravelWayfinderFireExitSignArrowPointDirectionSignal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 2.53a.34.34 0 0 0-.34 0a.22.22 0 0 0 0 .29c1 1.74 1.26 4.21-.17 5.56A4.69 4.69 0 0 1 2.2 6.73a3.3 3.3 0 0 0-1.7 3a4 4 0 0 0 4.25 3.77A3.88 3.88 0 0 0 9 9.69C9.11 7.16 7.33 4 4 2.53Z"/><path d="M6.87 9.69a1.69 1.69 0 0 1-1.7 1.69M11.5.5l2 2l-2 2m2-2H9"/></g>`), g.Group(children),
	)
}

func TravelWayfinderLadderBusinessProductMetaphorLadder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5.5v13m7-13v13M3.5 3h7m-7 4h7m-7 4h7"/>`), g.Group(children),
	)
}

func TravelWayfinderLifebuoyWaterLifeRingWheelyLifebeltKisbee(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 1.5L9.09 4.91M1.5 1.5l3.41 3.41M1.5 12.5l3.41-3.41m7.59 3.41L9.09 9.09"/><circle cx="7" cy="7" r="6.5"/><circle cx="7" cy="7" r="2.95"/></g>`), g.Group(children),
	)
}

func TravelWayfinderLiftArrowUpHumanDownPersonUserLiftElevator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="2.5" r="2"/><path d="M1.5 6.5h2a1 1 0 0 1 1 1v3h0h-4h0v-3a1 1 0 0 1 1-1Zm-1 4v3m4-3v3m4-8L11 3l2.5 2.5m-5 4L11 12l2.5-2.5"/></g>`), g.Group(children),
	)
}

func TravelWayfinderManArmRaisesOneManRaiseArmScaningDetectPostureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M2.5 6.5h9m-3 0v4h-3v-4m0 4v3m3-3v3"/></g>`), g.Group(children),
	)
}

func TravelWayfinderManArmRaisesTwoManRaiseArmScaningDetectPostureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M2.5 5a7.5 7.5 0 0 0 9 0"/><path d="M8.5 6.33v4.17h-3V6.33m0 4.17v3m3-3v3"/></g>`), g.Group(children),
	)
}

func TravelWayfinderManSymbolGeometricGenderBoyPersonMaleHumanUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.5" cy="9.5" r="4"/><path d="M9 .5h4.5V5M7.33 6.67L13.5.5"/></g>`), g.Group(children),
	)
}

func TravelWayfinderStairsOneStairsStaircaseRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5H9V5H5v4H.5v4.5h13V.5z"/>`), g.Group(children),
	)
}

func TravelWayfinderStairsTwoStairsStaircaseLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5H5V5h4v4h4.5v4.5H.5V.5z"/>`), g.Group(children),
	)
}

func TravelWayfinderToiletSignManToiletSignRestroomBathroomUserHumanPersonManMale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M7 13.5c-3 0-3-7-3-7h6s0 7-3 7Z"/></g>`), g.Group(children),
	)
}

func TravelWayfinderToiletSignManWomanToiletSignRestroomBathroomUserHumanPerson(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="2.5" r="2"/><path d="M11 6.5c-2.5 0-2.5 7-2.5 7h5s0-7-2.5-7Z"/><circle cx="3" cy="2.5" r="2"/><path d="M3 13.5c-2.5 0-2.5-7-2.5-7h5s0 7-2.5 7Z"/></g>`), g.Group(children),
	)
}

func TravelWayfinderToiltetSignWomanToiletSignRestroomBathroomUserHumanPersonWomanFemale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M7 6.5c-3 0-3 7-3 7h6s0-7-3-7Z"/></g>`), g.Group(children),
	)
}

func TravelWayfinderWomanSymbolGeometricGenderFemalePersonHumanUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="4" r="3.5"/><path d="M7 7.5v6M5 11h4"/></g>`), g.Group(children),
	)
}
