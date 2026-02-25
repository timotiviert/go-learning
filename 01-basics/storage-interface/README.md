// Erstelle ein Interface Storage mit:
// - Save(key, value string) error
// - Load(key string) (string, error)
//
// Implementiere:
// - FileStorage (speichert in Dateien)
// - MemoryStorage (speichert in Map)
// - Bonus: PostgresStorage (einfache DB connection)
//
// Main: Nutze beide über das Interface polymorphisch
