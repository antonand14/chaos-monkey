import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [experiments, setExperiments] = useState([])
  const [newExperiment, setNewExperiment] = useState({
    name: '',
    type: 'pod_assassin',
    target_namespace: 'default',
    target_labels: {},
    blast_radius: 50,
    duration_seconds: 60,
    dry_run: false
  })

  useEffect(() => {
    fetchExperiments()
  }, [])

  const fetchExperiments = async () => {
    try {
      const response = await fetch('/api/experiments')
      const data = await response.json()
      setExperiments(data)
    } catch (error) {
      console.error('Error fetching experiments:', error)
    }
  }

  const createExperiment = async () => {
    try {
      const response = await fetch('/api/experiments', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(newExperiment)
      })
      const data = await response.json()
      console.log('Created experiment:', data)
      fetchExperiments()
    } catch (error) {
      console.error('Error creating experiment:', error)
    }
  }

  const runExperiment = async (experimentId) => {
    try {
      await fetch(`/api/experiments/${experimentId}/run`, { method: 'POST' })
      console.log('Running experiment:', experimentId)
    } catch (error) {
      console.error('Error running experiment:', error)
    }
  }

  return (
    <div className="App">
      <header>
        <h1>🐵 Chaos Monkey as a Service</h1>
        <p>Test your cluster resilience through controlled chaos</p>
      </header>

      <main>
        <section className="create-experiment">
          <h2>Create Chaos Experiment</h2>
          <div className="form">
            <input
              type="text"
              placeholder="Experiment name"
              value={newExperiment.name}
              onChange={(e) => setNewExperiment({...newExperiment, name: e.target.value})}
            />
            <select
              value={newExperiment.type}
              onChange={(e) => setNewExperiment({...newExperiment, type: e.target.value})}
            >
              <option value="pod_assassin">Pod Assassin</option>
              <option value="network_gremlin">Network Gremlin</option>
              <option value="resource_hog">Resource Hog</option>
            </select>
            <input
              type="text"
              placeholder="Namespace"
              value={newExperiment.target_namespace}
              onChange={(e) => setNewExperiment({...newExperiment, target_namespace: e.target.value})}
            />
            <input
              type="number"
              placeholder="Blast radius %"
              value={newExperiment.blast_radius}
              onChange={(e) => setNewExperiment({...newExperiment, blast_radius: parseInt(e.target.value)})}
            />
            <button onClick={createExperiment}>Create Experiment</button>
          </div>
        </section>

        <section className="experiments-list">
          <h2>Experiments</h2>
          {experiments.length === 0 ? (
            <p>No experiments yet. Create one above!</p>
          ) : (
            <div className="experiments">
              {experiments.map((exp) => (
                <div key={exp.id} className="experiment-card">
                  <h3>{exp.name}</h3>
                  <p>Type: {exp.type}</p>
                  <p>Namespace: {exp.target_namespace}</p>
                  <p>Blast Radius: {exp.blast_radius}%</p>
                  <button onClick={() => runExperiment(exp.id)}>▶ Run</button>
                </div>
              ))}
            </div>
          )}
        </section>

        <section className="leaderboard">
          <h2>🏆 Resilience Leaderboard</h2>
          <p>Coming soon...</p>
        </section>
      </main>
    </div>
  )
}

export default App
