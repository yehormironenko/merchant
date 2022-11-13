package config

import "time"

type Server struct {
	Port            string        `koanf:"port"`
	ShutdownTimeout time.Duration `koanf:"shutdownTimeout"`
}

/*func (s Server) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("port", s.Port)
	enc.AddDuration("shutdownTimeout", s.ShutdownTimeout)
	return enc.AddObject("kubernetesProbes", s.KubernetesProbes)
}*/
