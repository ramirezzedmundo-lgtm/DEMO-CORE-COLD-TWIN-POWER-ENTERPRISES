import requests
import time
import random

URL = "http://localhost:8080/v1/ingest"

print("🚀 Iniciando Recuperación de Capital...")
try:
    while True:
        # Simulamos que la IA salvó un producto
        res = requests.post(URL)
        if res.status_code == 200:
            print("✅ Dinero recuperado enviado al Core.")
        time.sleep(1) # Un evento por segundo para que el cliente vea el flujo
except:
    print("🛑 Demo finalizado.")
