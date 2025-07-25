package rpicamera

type params struct {
	LogLevel              string
	CameraID              uint32
	Width                 uint32
	Height                uint32
	HFlip                 bool
	VFlip                 bool
	Brightness            float32
	Contrast              float32
	Saturation            float32
	Sharpness             float32
	Exposure              string
	AWB                   string
	AWBGainRed            float32
	AWBGainBlue           float32
	Denoise               string
	Shutter               uint32
	Metering              string
	Gain                  float32
	EV                    float32
	ROI                   string
	HDR                   bool
	TuningFile            string
	Mode                  string
	FPS                   float32
	AfMode                string
	AfRange               string
	AfSpeed               string
	LensPosition          float32
	AfWindow              string
	FlickerPeriod         uint32
	TextOverlayEnable     bool
	TextOverlay           string
	Codec                 string
	IDRPeriod             uint32
	Bitrate               uint32
	HardwareH264Profile   string
	HardwareH264Level     string
	SoftwareH264Profile   string
	SoftwareH264Level     string
	SecondaryWidth        uint32
	SecondaryHeight       uint32
	SecondaryFPS          float32
	SecondaryMJPEGQuality uint32
}
