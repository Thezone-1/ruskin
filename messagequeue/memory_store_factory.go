package messagequeue

type MemoryStoreFactoryConfig struct {
}

type MemoryStoreFactory struct {
	*StoreFactoryConfig
}

func NewMemoryStoreFactory(config *StoreFactoryConfig) *MemoryStoreFactory {
	return &MemoryStoreFactory{
		StoreFactoryConfig: config,
	}
}

func (msf *MemoryStoreFactory) Produce(topicName string) Store {

	return NewMemoryStore(topicName)

}
