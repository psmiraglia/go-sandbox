package sandbox

// DoIt executes the task related to "doit" command.
func DoIt() {
	log := Logger
	log.Debug("Hello! This is Debug...")
	log.Info("Hello! This is Info...")
	log.Warning("Hello! This is Warning...")
	log.Error("Hello! This is Error...")
	log.Fatal("Hello! This is Fatal...")
	log.Panic("Hello! This is Panic...")
}
